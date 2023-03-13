import React, {useEffect} from 'react';
import {StyleSheet} from 'react-native';
import {NavigationContainer} from '@react-navigation/native';
//import {Platform} from 'react-native';

import {createNativeStackNavigator} from '@react-navigation/native-stack';
import Home from './pages/Home';
import Pair from './pages/Pair';

// Import polyfills if not running on web.  Attempting to import these in web mode will result in numerous errors
// trying to access react-native APIs
/* if (Platform.OS !== 'web') {
  // @ts-expect-error
  import('react-native-polyfill-globals');
} */

function App() {
  const [paired, setPaired] = React.useState(true);

  const Stack = createNativeStackNavigator();

  //custom transition that moves to the left

  //The pairing authentication is actually not related to the bluetooth logic
  //So all of the bluetooth logic can be

  return (
    <NavigationContainer>
      <Stack.Navigator
        screenOptions={{headerShown: false}}
        initialRouteName="Home">
        <Stack.Screen
          name="Pair"
          component={Pair}
          initialParams={{Stack, styles}}
        />
        <Stack.Screen
          name="Home"
          component={Home}
          initialParams={{Stack, styles}}
        />
      </Stack.Navigator>
    </NavigationContainer>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: 'white',
  },
});

export default App;
