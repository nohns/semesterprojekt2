import React, {useState, useEffect} from 'react';

import {
  View,
  StyleSheet,
  Text,
  useWindowDimensions,
  FlatList,
  Animated,
  ActivityIndicator,
  TouchableOpacity,
  Alert,
} from 'react-native';
import {CheckIcon, ChevronRightIcon} from 'react-native-heroicons/outline';

interface OnboardingProps {}

function Onboarding({}: OnboardingProps) {
  const [isLoading, setIsloading] = useState(true);
  const [isConnecting, setIsConnecting] = useState(false);
  const [isConnected, setIsConnected] = useState(false);

  useEffect(() => {
    setTimeout(() => {
      setIsloading(false);
    }, 2000);
  }, []);

  function onConnect() {
    setIsConnecting(true);
    setTimeout(() => {
      setIsConnecting(false);
      setIsConnected(true);
      Alert.alert(
        'Forbundet',
        'Du er nu forbundet til din bridge med pin 1234',
      );
    }, 2500);
  }

  return (
    <View style={[styles.container]}>
      <Text style={styles.title}>Find din bridge</Text>
      {isLoading && (
        <View style={styles.loadView}>
          <View style={styles.loader}>
            <ActivityIndicator animating={true} />
            <Text style={styles.loaderText}>Søger i nærheden</Text>
          </View>
        </View>
      )}
      {!isLoading && (
        <View style={styles.itemContainer}>
          <TouchableOpacity style={styles.item} onPress={onConnect}>
            <Text style={styles.itemText}>Smart Lock Bridge</Text>
            <View style={styles.itemAction}>
              <Text style={styles.itemActionText}>
                {isConnecting
                  ? 'Forbinder...'
                  : isConnected
                  ? 'Forbundet'
                  : 'Forbind'}
              </Text>
              {isConnecting ? (
                <ActivityIndicator animating={true} />
              ) : isConnected ? (
                <CheckIcon size={24} color={'black'} />
              ) : (
                <ChevronRightIcon size={24} color={'black'} />
              )}
            </View>
          </TouchableOpacity>
        </View>
      )}
    </View>
  );
}

export default Onboarding;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    paddingTop: 60,
  },
  itemContainer: {
    width: '100%',
    paddingHorizontal: 20,
    paddingTop: 20,
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
    marginLeft: 20,
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
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    paddingBottom: 100,
  },
});
