import Crypto from 'react-native-crypto';
//Still needs to be installed //TODO:

const generateCSR = async () => {
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
