import Crypto from 'react-native-quick-crypto';
import {Buffer} from '@craftzdog/react-native-buffer';
//Still needs to be installed //TODO:

const generateCSR = async () => {
  const {pubKey, privKey} = await generateKeyPair();
  
  Crypto.X509Certificat
    const keyPair = await Crypto.RSA.generateKeyPair(2048);
    const csr = Crypto.X509.createCSR({
      subject: {
        CommonName: 'localhost',
        Country: 'DK',
        Province: 'EU',
        Locality: 'Copenhagen',
        Organization: 'Dev',
        OrganizationalUnit: 'Semesterprojekt',
      },
      key: keyPair.privateKey,
    });


  // Transfer the CSR to the Go server for signing...
  // ...
  return csr;
};

const generateKeyPair = () => {
  return new Promise<{
    pubKey?: Buffer;
    privKey?: Buffer;
  }>((resolve, reject) => {
    Crypto.generateKeyPair(
      'RSA',
      {
        modulusLength: 4096,
        publicKeyEncoding: {
          type: 'spki',
          format: 'pem',
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
}