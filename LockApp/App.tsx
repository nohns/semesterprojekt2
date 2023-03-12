import React from 'react';
import {StyleSheet} from 'react-native';
import {NavigationContainer} from '@react-navigation/native';

import {createNativeStackNavigator} from '@react-navigation/native-stack';
import Home from './pages/Home';
import Pair from './pages/Pair';

function App() {
  const [paired, setPaired] = React.useState(true);

  const Stack = createNativeStackNavigator();

  //custom transition that moves to the left

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
