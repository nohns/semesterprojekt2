import React from 'react';
import {View, StyleSheet, Dimensions} from 'react-native';
import Animated from 'react-native-reanimated';
import {
  useDerivedValue,
  useAnimatedStyle,
  withTiming,
  interpolateColor,
} from 'react-native-reanimated';

import {LockClosedIcon} from 'react-native-heroicons/outline';
import {LockOpenIcon} from 'react-native-heroicons/outline/';

interface LockedProps {
  locked: boolean;
}

const Colors = {
  green: {
    background: '#1E1E1E',
    circle: 'green',
    text: '#F8F8F8',
  },
  red: {
    background: '#F8F8F8',
    circle: 'red',
    text: '#1E1E1E',
  },
};

function Locked({locked}: LockedProps) {
  const progress = useDerivedValue(() => {
    return locked ? withTiming(1) : withTiming(0);
  }, [locked]);

  const rCircleStyle = useAnimatedStyle(() => {
    const backgroundColor = interpolateColor(
      progress.value,
      [0, 1],
      [Colors.green.circle, Colors.red.circle],
    );
    return {
      backgroundColor,
    };
  });

  return (
    <Animated.View style={[styles.container]}>
      <Animated.View style={[styles.smallCircle, rCircleStyle]}>
        <Animated.View style={[styles.circle, rCircleStyle]}>
          {locked ? (
            <LockClosedIcon size={65} color={'white'} style={styles.lock} />
          ) : (
            <LockOpenIcon
              size={65}
              color={'white'}
              opacity={'100'}
              style={styles.lock}
            />
          )}
        </Animated.View>
      </Animated.View>
    </Animated.View>
  );
}
const SIZE = Dimensions.get('window').width * 0.6;
const SMALLSIZE = Dimensions.get('window').width * 0.5;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
  circle: {
    width: SIZE,
    height: SIZE,
    borderRadius: SIZE / 2,
    justifyContent: 'center',
    alignItems: 'center',
    shadowOffset: {
      width: 0,
      height: 10,
    },
    shadowRadius: 20,
    shadowOpacity: 0.8,
    opacity: 0.8,
  },
  smallCircle: {
    width: SMALLSIZE,
    height: SMALLSIZE,
    borderRadius: SMALLSIZE / 2,
    justifyContent: 'center',
    alignItems: 'center',
    shadowOffset: {
      width: 0,
      height: 10,
    },
    shadowRadius: 10,
    shadowOpacity: 0.2,
    opacity: 0.3,
  },
  lock: {
    justifyContent: 'center',
    alignItems: 'center',
  },
});

export default Locked;
