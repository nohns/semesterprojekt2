import React, {useState, useEffect, useCallback, useRef} from 'react';

import {
  View,
  StyleSheet,
  Text,
  FlatList,
  ActivityIndicator,
  TouchableOpacity,
  Alert,
  NativeModules,
  NativeEventEmitter,
} from 'react-native';
import {
  ArrowPathIcon,
  CheckIcon,
  ChevronRightIcon,
} from 'react-native-heroicons/outline';
import {Buffer} from '@craftzdog/react-native-buffer';

import BleManager, {
  BleDisconnectPeripheralEvent,
  BleManagerDidUpdateValueForCharacteristicEvent,
  Peripheral,
} from 'react-native-ble-manager';
import {sleep} from '../util';
import {Colors} from 'react-native/Libraries/NewAppScreen';
import {Identity} from '../model/identity';
import {
  ALLOW_DUPLICATES,
  SECONDS_TO_SCAN_FOR,
  SERVICE_UUIDS,
  SMART_LOCK_CHARACTERISTIC_CERT_LEN_UUID,
  SMART_LOCK_CHARACTERISTIC_CERT_UUID,
  SMART_LOCK_SERVICE_ID,
} from '../bluetooth/constants';
import {startIdentityExchange} from '../bluetooth/exchange';

const BleManagerModule = NativeModules.BleManager;
const bleManagerEmitter = new NativeEventEmitter(BleManagerModule);

declare module 'react-native-ble-manager' {
  // enrich local contract with custom state properties needed by App.tsx
  interface Peripheral {
    connected?: boolean;
    connecting?: boolean;
  }
}

interface OnboardingProps {
  onIdentityGiven: (identity: Identity) => void;
}

function Onboarding({onIdentityGiven}: OnboardingProps) {
  const [isScanning, setIsScanning] = useState(false);
  const [peripherals, setPeripherals] = useState<
    Record<Peripheral['id'], Peripheral>
  >({});

  const addOrUpdatePeripheral = (id: string, updatedPeripheral: Peripheral) => {
    setPeripherals(old => ({
      ...old,
      [id]: {
        ...(old[id] ?? []),
        ...updatedPeripheral,
      },
    }));
  };

  const startScan = useCallback(() => {
    try {
      console.debug('[startScan] starting scan...');
      setIsScanning(true);
      BleManager.scan(SERVICE_UUIDS, SECONDS_TO_SCAN_FOR, ALLOW_DUPLICATES, {})
        .then(() => {
          console.debug('[startScan] scan promise returned successfully.');
        })
        .catch(err => {
          console.error('[startScan] ble scan returned in error', err);
        });
    } catch (error) {
      console.error('[startScan] ble scan error thrown', error);
    }
  }, []);

  const handleStopScan = () => {
    setIsScanning(false);
    console.debug('[handleStopScan] scan is stopped.');
  };

  const handleDisconnectedPeripheral = (
    event: BleDisconnectPeripheralEvent,
  ) => {
    let peripheral = peripherals[event.peripheral];
    if (peripheral) {
      console.debug(
        `[handleDisconnectedPeripheral][${peripheral.id}] previously connected peripheral is disconnected.`,
        event.peripheral,
      );
      addOrUpdatePeripheral(peripheral.id, {...peripheral, connected: false});
    }
    console.debug(
      `[handleDisconnectedPeripheral][${event.peripheral}] disconnected.`,
    );
  };

  const certBufferRef = useRef<Buffer | undefined>();
  const certBufferBytesReceivedRef = useRef<number>(0);
  const handleUpdateValueForCharacteristic = (
    data: BleManagerDidUpdateValueForCharacteristicEvent,
  ) => {
    switch (data.characteristic) {
      case SMART_LOCK_CHARACTERISTIC_CERT_LEN_UUID:
        // Convert 16-bit number (in two bytes) to a js number.
        // eslint-disable-next-line no-bitwise
        const readLen = data.value[0] | (data.value[1] << 8);
        certBufferRef.current = Buffer.alloc(readLen, 0);
        certBufferBytesReceivedRef.current = 0;
        console.debug('[handleUpdateValueForCharacteristic] readLen=', readLen);
        break;

      case SMART_LOCK_CHARACTERISTIC_CERT_UUID:
        if (!certBufferRef.current) {
          break;
        }

        // Write data to buffer holding the received cert
        certBufferRef.current.set(
          data.value,
          certBufferBytesReceivedRef.current,
        );
        certBufferBytesReceivedRef.current += data.value.length;

        if (
          certBufferBytesReceivedRef.current === certBufferRef.current.length
        ) {
          // Done reading cert

          addOrUpdatePeripheral(data.peripheral, {
            connecting: false,
            connected: true,
          } as any);

          Alert.alert(
            'Bridge forbudnet',
            'Din pinkode er til din Smart Lock er 1234',
            [
              {
                text: 'OK',
                onPress: async () => {
                  if (!certBufferRef.current) {
                    return;
                  }
                  await sleep(1000);
                  // Set identity
                  onIdentityGiven({
                    base64Certificate: certBufferRef.current.toString('base64'),
                  });
                },
              },
            ],
          );
        }
        break;
    }
  };

  const handleDiscoverPeripheral = (peripheral: Peripheral) => {
    console.debug('[handleDiscoverPeripheral] new BLE peripheral=', peripheral);
    if (!peripheral.name) {
      peripheral.name = 'NO NAME';
    }
    addOrUpdatePeripheral(peripheral.id, peripheral);
  };

  const connectPeripheral = async (peripheral: Peripheral) => {
    if (peripheral.connected) {
      return;
    }

    try {
      if (peripheral) {
        addOrUpdatePeripheral(peripheral.id, {...peripheral, connecting: true});

        await BleManager.connect(peripheral.id);
        console.debug(`[connectPeripheral][${peripheral.id}] connected.`);

        // before retrieving services, it is often a good idea to let bonding & connection finish properly
        await sleep(900);

        /* Test read current RSSI value, retrieve services first */
        const peripheralData = await BleManager.retrieveServices(peripheral.id);
        console.debug(
          `[connectPeripheral][${peripheral.id}] retrieved peripheral services`,
          peripheralData,
        );

        const rssi = await BleManager.readRSSI(peripheral.id);
        console.debug(
          `[connectPeripheral][${peripheral.id}] retrieved current RSSI value: ${rssi}.`,
        );

        let p = peripherals[peripheral.id];
        if (p) {
          addOrUpdatePeripheral(peripheral.id, {...peripheral, rssi});
        }

        await BleManager.startNotification(
          peripheral.id,
          SMART_LOCK_SERVICE_ID,
          SMART_LOCK_CHARACTERISTIC_CERT_LEN_UUID,
        );
        console.log('Notification for cert len started');

        await BleManager.startNotification(
          peripheral.id,
          SMART_LOCK_SERVICE_ID,
          SMART_LOCK_CHARACTERISTIC_CERT_UUID,
        );
        console.log('Notification for cert data started');

        await startIdentityExchange(peripheral);
      }
    } catch (error) {
      console.error(
        `[connectPeripheral][${peripheral.id}] connectPeripheral error`,
        error,
      );
    }
  };

  // Start Bluetooth, attach listeners and start scanning for bridges
  useEffect(() => {
    try {
      BleManager.start({showAlert: false})
        .then(() => {
          console.debug('BleManager started.');
          startScan();
        })
        .catch(error =>
          console.error('BeManager could not be started.', error),
        );
    } catch (error) {
      console.error('unexpected error starting BleManager.', error);
      return;
    }

    const listeners = [
      bleManagerEmitter.addListener(
        'BleManagerDiscoverPeripheral',
        handleDiscoverPeripheral,
      ),
      bleManagerEmitter.addListener('BleManagerStopScan', handleStopScan),
      bleManagerEmitter.addListener(
        'BleManagerDisconnectPeripheral',
        handleDisconnectedPeripheral,
      ),
      bleManagerEmitter.addListener(
        'BleManagerDidUpdateValueForCharacteristic',
        handleUpdateValueForCharacteristic,
      ),
    ];

    return () => {
      console.debug('[app] main component unmounting. Removing listeners...');
      for (const listener of listeners) {
        listener.remove();
      }
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const renderItem = ({item}: {item: Peripheral}) => {
    return (
      <TouchableOpacity
        onPress={() => connectPeripheral(item)}
        style={styles.item}>
        <Text style={styles.itemText}>{item.advertising.localName}</Text>
        <View style={styles.itemAction}>
          <Text style={styles.itemActionText}>
            {item.connecting
              ? 'Forbinder...'
              : item.connected
              ? 'Forbundet'
              : 'Forbind'}
          </Text>
          {item.connecting ? (
            <ActivityIndicator animating={true} />
          ) : item.connected ? (
            <CheckIcon size={24} color={'black'} />
          ) : (
            <ChevronRightIcon size={24} color={'black'} />
          )}
        </View>
      </TouchableOpacity>
    );
  };

  const peripheralsArray = Object.values(peripherals);

  return (
    <View style={[styles.container]}>
      <Text style={styles.title}>Find din bridge</Text>

      {peripheralsArray.length > 0 && (
        <View style={styles.itemContainer}>
          <FlatList
            data={peripheralsArray}
            contentContainerStyle={{rowGap: 12}}
            renderItem={renderItem}
            keyExtractor={item => item.id}
          />
        </View>
      )}

      <View
        style={[
          styles.loadView,
          {
            flex: peripheralsArray.length ? undefined : 1,
            paddingTop: peripheralsArray.length ? 20 : 0,
          },
        ]}>
        {(peripheralsArray.length === 0 || isScanning) && (
          <View style={styles.loader}>
            {isScanning && <ActivityIndicator animating={true} />}

            <Text style={styles.loaderText}>
              {peripheralsArray.length === 0 &&
                !isScanning &&
                'Ingen enheder fundet'}
              {isScanning && 'Søger i nærheden'}
            </Text>
          </View>
        )}
        {!isScanning && (
          <TouchableOpacity onPress={startScan}>
            <View style={styles.button}>
              <ArrowPathIcon size={24} color={'#222'} />
              <Text style={styles.buttonText}>Søg igen</Text>
            </View>
          </TouchableOpacity>
        )}
      </View>
    </View>
  );
}

export default Onboarding;

const boxShadow = {
  shadowColor: '#000',
  shadowOffset: {
    width: 0,
    height: 2,
  },
  shadowOpacity: 0.25,
  shadowRadius: 3.84,
  elevation: 5,
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'flex-start',
    backgroundColor: '#eee',
    paddingTop: 60,
    paddingHorizontal: 20,
  },
  itemContainer: {
    width: '100%',
    paddingTop: 20,
  },
  button: {
    height: 40,
    backgroundColor: '#ffffff',
    borderRadius: 20,
    paddingVertical: 10,
    gap: 5,
    alignItems: 'center',
    flexDirection: 'row',
    paddingHorizontal: 15,
    marginBottom: 20,
  },
  buttonText: {
    color: '#222',
    fontWeight: 'bold',
  },
  item: {
    width: '100%',
    height: 50,
    backgroundColor: 'white',
    borderRadius: 20,
    padding: 15,
    justifyContent: 'space-between',
    alignItems: 'center',
    flexDirection: 'row',
  },
  itemText: {
    fontWeight: 'bold',
  },
  itemAction: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 8,
  },
  itemActionText: {
    //fontWeight: 'bold',
    color: '#aaa',
  },
  title: {
    fontSize: 40,
    marginBottom: 5,
    marginTop: 20,
  },
  description: {
    fontWeight: '300',
    color: 'red',
    textAlign: 'center',
    paddingHorizontal: 64,
  },
  loaderText: {
    fontSize: 20,
  },
  loader: {
    width: '100%',
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    gap: 5,
  },
  loadView: {
    width: '100%',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 20,
    paddingBottom: 100,
  },
  /*
  dfdsfdsf
  */
  engine: {
    position: 'absolute',
    right: 10,
    bottom: 0,
    color: Colors.black,
  },
  scanButton: {
    alignItems: 'center',
    justifyContent: 'center',
    paddingVertical: 16,
    backgroundColor: '#0a398a',
    margin: 10,
    borderRadius: 12,
    ...boxShadow,
  },
  scanButtonText: {
    fontSize: 20,
    letterSpacing: 0.25,
    color: Colors.white,
  },
  body: {
    backgroundColor: '#0082FC',
    flex: 1,
  },
  sectionContainer: {
    marginTop: 32,
    paddingHorizontal: 24,
  },
  sectionTitle: {
    fontSize: 24,
    fontWeight: '600',
    color: Colors.black,
  },
  sectionDescription: {
    marginTop: 8,
    fontSize: 18,
    fontWeight: '400',
    color: Colors.dark,
  },
  highlight: {
    fontWeight: '700',
  },
  footer: {
    color: Colors.dark,
    fontSize: 12,
    fontWeight: '600',
    padding: 4,
    paddingRight: 12,
    textAlign: 'right',
  },
  peripheralName: {
    fontSize: 16,
    textAlign: 'center',
    padding: 10,
  },
  rssi: {
    fontSize: 12,
    textAlign: 'center',
    padding: 2,
  },
  peripheralId: {
    fontSize: 12,
    textAlign: 'center',
    padding: 2,
    paddingBottom: 20,
  },
  row: {
    marginLeft: 10,
    marginRight: 10,
    borderRadius: 20,
    ...boxShadow,
  },
  noPeripherals: {
    margin: 10,
    textAlign: 'center',
    color: Colors.white,
  },
});
