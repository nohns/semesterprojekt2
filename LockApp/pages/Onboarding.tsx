import React from 'react';

import {
  View,
  StyleSheet,
  Text,
  useWindowDimensions,
  FlatList,
  Animated,
} from 'react-native';

interface OnboardingProps {
  navigation: any;
}

function Onboarding({navigation}: OnboardingProps) {
  const [currentIndex, setCurrentIndex] = React.useState(0);
  const scrollX = React.useRef(new Animated.Value(0)).current;
  const slidesRef = React.useRef(null);

  const viewConfig = React.useRef({
    viewAreaCoveragePercentThreshold: 50,
  }).current;

  const viewableItemsChanged = React.useRef(({viewableItems}: any) => {
    setCurrentIndex(viewableItems[0].index);
  }).current;

  return (
    <View style={[styles.container]}>
      <View style={{flex: 3}}>
        <FlatList
          data={slides}
          renderItem={({item}) => <OnboardingItem item={item} />}
          horizontal
          showsHorizontalScrollIndicator
          pagingEnabled
          bounces={false}
          onScroll={Animated.event(
            [{nativeEvent: {contentOffset: {x: scrollX}}}],
            {useNativeDriver: false},
          )}
          scrollEventThrottle={32}
          onViewableItemsChanged={viewableItemsChanged}
          viewabilityConfig={viewConfig}
          ref={slidesRef}
        />
      </View>
    </View>
  );
}

export default Onboarding;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  title: {
    fontWeight: '800',
    fontSize: 28,
    marginBottom: 10,
    color: 'red',
    textAlign: 'center',
  },
  description: {
    fontWeight: '300',
    color: 'red',
    textAlign: 'center',
    paddingHorizontal: 64,
  },
});

const slides = [
  {
    id: '1',
    title: 'Title 1',
    description: 'Description 1',
  },
  {
    id: '2',
    title: 'Title 2',
    description: 'Description 2',
  },
  {
    id: '3',
    title: 'Title 3',
    description: 'Description 3',
  },
];

function OnboardingItem({item}: {item: {title: string; description: string}}) {
  const {width} = useWindowDimensions();
  return (
    <View style={{width}}>
      {/*  <View style={{flex: 0.7, justifyContent: 'center'}} />
        <View style={{flex: 0.3}}> */}
      <Text style={styles.title}>{item.title}</Text>
      <Text style={styles.description}>{item.description}</Text>
      {/*  </View> */}
    </View>
  );
}

function Paginator() {}
