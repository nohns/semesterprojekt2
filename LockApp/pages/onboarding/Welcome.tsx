import React, {useEffect} from 'react';
import {View, StyleSheet, StyleProp, ViewStyle, Text} from 'react-native';

interface WelcomeProps {
  style?: StyleProp<ViewStyle>;
}

function Welcome({style}: WelcomeProps) {
  useEffect(() => {}, []);

  return (
    <View style={[styles.container, style]}>
      <Text>Opsætning</Text>
      <Text>Din</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
});

export default Welcome;
