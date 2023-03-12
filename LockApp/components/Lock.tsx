import React from 'react';
import {View, StyleSheet} from 'react-native';
import {LockClosedIcon} from 'react-native-heroicons/outline';
import {LockOpenIcon} from 'react-native-heroicons/outline/';

interface LockProps {
  locked: boolean;
}

function Lock({locked}: LockProps) {
  return (
    <View style={[styles.container]}>
      <View style={[styles.lock]} />
      <View style={locked ? [styles.backgroundRed] : [styles.backgroundGreen]}>
        <View
          style={locked ? [styles.backgroundRed2] : [styles.backgroundGreen2]}
        />
        {locked ? (
          <LockClosedIcon size={65} color={'white'} style={styles.lock} />
        ) : (
          <LockOpenIcon size={65} color={'white'} style={styles.lock} />
        )}
      </View>
      <View />
    </View>
  );
}

export default Lock;

//style sheet for lock
const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'white',
  },
  backgroundRed: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'red',
    width: 175,
    height: 175,
    borderRadius: 100,
    opacity: 0.7,
  },
  backgroundRed2: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'red',
    width: 215,
    height: 215,
    borderRadius: 100,
    opacity: 0.4,
  },

  backgroundGreen: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'green',
    width: 175,
    height: 175,
    borderRadius: 100,
    opacity: 0.7,
  },
  backgroundGreen2: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'green',
    width: 215,
    height: 215,
    borderRadius: 100,
    opacity: 0.4,
  },
  lock: {
    position: 'absolute',
    justifyContent: 'center',
    alignItems: 'center',
  },
});

/* border-radius: 50%;
border: 2px solid red;
padding: 5px;
font-size: 30px; */
