import React from 'react';
import {View, StyleSheet, StyleProp, ViewStyle} from 'react-native';

interface PairProps {
  style?: StyleProp<ViewStyle>;
  navigation: any;
}

function Pair({navigation, style}: PairProps) {
  return <View style={[styles.container, style]}></View>;
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
});

export default Pair;
