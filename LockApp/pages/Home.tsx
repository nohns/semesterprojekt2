import React from 'react';
import {
  View,
  StyleSheet,
  StyleProp,
  ViewStyle,
  Text,
  TouchableWithoutFeedback,
} from 'react-native';
import {ChevronLeftIcon} from 'react-native-heroicons/outline';
import Locked from '../components/Locked';

interface HomeProps {
  navigation: any;
  style?: StyleProp<ViewStyle>;
}

function Home({navigation, style}: HomeProps) {
  const [locked, setLocked] = React.useState(true);

  const handleOnClick = () => {
    setLocked(!locked);

    fetch('http://localhost:8500/lock.v1/setLock', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        id: '123',
        engaged: !locked,
      }),
    });
  };

  const handleNavigate = () => {
    navigation.navigate('Onboarding');
  };

  /*   const SWITCH_TRACK_COLOR = {
    false: '#767577',
    true: '#81b0ff',
  }; */

  return (
    <View style={[styles.container, style]}>
      <TouchableWithoutFeedback onPress={handleNavigate}>
        <View style={styles.tilbage}>
          <ChevronLeftIcon size={30} color={'black'} />
          <Text>Tilbage </Text>
        </View>
      </TouchableWithoutFeedback>

      <Text style={styles.h1}> {'Smart Lock'}</Text>

      <View style={styles.circleWrapper}>
        <View style={locked ? [styles.circleRed] : [styles.circleGreen]} />

        {locked ? (
          <Text style={styles.h3}> {'Låsen er slået til'}</Text>
        ) : (
          <Text style={styles.h3}> {'Låsen er åben'}</Text>
        )}
      </View>

      <TouchableWithoutFeedback onPress={handleOnClick}>
        <View style={styles.bigWrapper}>
          {locked ? (
            <Text style={styles.h2}> {'Tryk for at låse'}</Text>
          ) : (
            <Text style={styles.h2}> {'Tryk for at låse op'}</Text>
          )}
          <Locked locked={locked} />
        </View>
      </TouchableWithoutFeedback>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    vh: 100,
    vw: 100,
    paddingTop: 60,
  },
  tilbage: {
    flexDirection: 'row',
    alignItems: 'center',

    fontSize: 20,
    marginTop: 60,
    marginLeft: 21,
    display: 'none',
  },
  circleWrapper: {
    marginLeft: 30,
    marginRight: 5,
    flexDirection: 'row',
    alignItems: 'center',
  },

  circleRed: {
    width: 20,
    height: 20,
    opacity: 0.6,
    borderRadius: 100,
    backgroundColor: 'red',
  },
  circleGreen: {
    width: 20,
    height: 20,
    opacity: 0.6,
    borderRadius: 100,
    backgroundColor: 'green',
  },

  h1: {
    fontSize: 40,
    marginLeft: 20,
    marginBottom: 5,
    marginTop: 20,
  },

  h2: {
    fontSize: 22,
    margin: 5,
    marginBottom: 40,
    marginTop: 125,
    padding: 75,
    opacity: 0.5,
  },

  h3: {
    fontSize: 15,

    opacity: 0.5,
  },
  bigWrapper: {
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    gap: 20,
  },
  bigContainer: {
    flexDirection: 'column',
  },
});

export default Home;
