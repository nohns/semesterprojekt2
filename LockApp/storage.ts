import {MMKVLoader} from 'react-native-mmkv-storage';

export const persistantStorage = new MMKVLoader().withEncryption().initialize();
