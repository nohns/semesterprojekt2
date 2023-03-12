import React from 'react';
import {View, StyleSheet, StyleProp, ViewStyle, Text} from 'react-native';

interface screenProps {
  style?: StyleProp<ViewStyle>;
}

function screen({style}: screenProps) {
  return (
    <View style={[styles.container, style]}>
      <Text>Hello World</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
});

export default screen;
