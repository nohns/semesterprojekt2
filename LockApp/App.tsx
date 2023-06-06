import React, {useEffect, useState} from 'react';
import {
  ActivityIndicator,
  SafeAreaView,
  StatusBar,
  StyleSheet,
  Text,
  View,
} from 'react-native';
import Onboarding from './pages/Onboarding';
import {Identity} from './model/identity';
import {persistantStorage} from './storage';
import Authenticated from './pages/Authenticated';

function App() {
  // App state
  const [identity, setIdentity] = useState<Identity | null>(null);
  const [state, setState] = useState<
    'onboarding' | 'authenticated' | 'initializing'
  >('initializing');

  // On mount check if we have an identity stored in persistent storage
  useEffect(() => {
    persistantStorage.getMapAsync<Identity>('identity').then(savedIdentity => {
      if (!savedIdentity) {
        setState('onboarding');
        return;
      }

      setIdentity(savedIdentity);
      setState('authenticated');
    });
  }, []);

  // When identity is given, store it in persistent storage and set state to authenticated
  async function onIdentityGiven(newIdentity: Identity) {
    setIdentity(newIdentity);
    await persistantStorage.setMapAsync('identity', newIdentity);
    setState('authenticated');
  }

  return (
    <>
      <StatusBar barStyle={'dark-content'} />
      {state === 'initializing' && (
        <SafeAreaView>
          <View style={styles.container}>
            <Text>Loading...</Text>
            <ActivityIndicator animating />
          </View>
        </SafeAreaView>
      )}
      {state === 'onboarding' && (
        <Onboarding onIdentityGiven={onIdentityGiven} />
      )}
      {state === 'authenticated' && identity && (
        <Authenticated identity={identity} />
      )}
    </>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: 'white',
    justifyContent: 'center',
    alignItems: 'center',
  },
});

export default App;
