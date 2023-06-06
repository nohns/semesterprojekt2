import {Buffer} from '@craftzdog/react-native-buffer';
import Crypto from 'react-native-quick-crypto';

export const generateKeyPair = () => {
  return new Promise<{
    pubKey?: Buffer;
    privKey?: Buffer;
  }>((resolve, reject) => {
    Crypto.generateKeyPair(
      'rsa',
      {
        modulusLength: 4096,
        publicKeyEncoding: {
          type: 'pkcs1',
          format: 'der',
        },
        privateKeyEncoding: {
          type: 'pkcs8',
          format: 'pem',
          cipher: 'aes-256-cbc',
          passphrase: 'top secret',
        },
      },
      (err, pubKey, privKey) => {
        if (err) {
          reject(err);
        }

        resolve({pubKey, privKey});
      },
    );
  });
};
