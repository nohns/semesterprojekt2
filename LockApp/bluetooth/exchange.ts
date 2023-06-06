import BleManager, {Peripheral} from 'react-native-ble-manager';
import {generateKeyPair} from '../csr/csr';
import {
  SMART_LOCK_CHARACTERISTIC_CSR_LEN_UUID,
  SMART_LOCK_CHARACTERISTIC_CSR_UUID,
  SMART_LOCK_SERVICE_ID,
} from './constants';
import {sleep} from '../util';

export async function startIdentityExchange(peripheral: Peripheral) {
  const {pubKey} = await generateKeyPair();

  const len = pubKey?.length;
  if (!len) {
    console.error('[startIdentityExchange] no pubKey length found.');
    return;
  }
  // Send public key length
  await write16BitValue(
    peripheral.id,
    SMART_LOCK_SERVICE_ID,
    SMART_LOCK_CHARACTERISTIC_CSR_LEN_UUID,
    len,
  );

  // Give time for the device to process the len
  await sleep(250);

  // Send public key
  try {
    const data = Buffer.from(pubKey).toJSON().data;

    await BleManager.writeWithoutResponse(
      peripheral.id,
      SMART_LOCK_SERVICE_ID,
      SMART_LOCK_CHARACTERISTIC_CSR_UUID,
      data,
      128,
      100,
    );
  } catch (err) {
    throw err;
  }
}

async function write16BitValue(
  peripheralId: string,
  serviceUUID: string,
  characteristicUUID: string,
  value: number,
) {
  try {
    const hexValue = value.toString(16).padStart(4, '0');
    await BleManager.writeWithoutResponse(
      peripheralId,
      serviceUUID,
      characteristicUUID,
      Buffer.from(hexValue, 'hex').toJSON().data,
      128,
    );
  } catch (err) {
    console.error(err);
  }
}
