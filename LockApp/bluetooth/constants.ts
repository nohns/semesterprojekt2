// Define UUID for the smart lock service we want to use of BLE peripheral
export const SMART_LOCK_SERVICE_ID = '9b7155fc-d47e-4309-9c81-a2261d582810';

// Define UUIDs for the smart lock characteristics we want to use of BLE peripheral
export const SMART_LOCK_CHARACTERISTIC_CSR_UUID =
  '9B7155FC-D47E-4309-9C81-A2261D582811';
export const SMART_LOCK_CHARACTERISTIC_CERT_UUID =
  '9B7155FC-D47E-4309-9C81-A2261D582812';
export const SMART_LOCK_CHARACTERISTIC_CSR_LEN_UUID =
  '9B7155FC-D47E-4309-9C81-A2261D582813';
export const SMART_LOCK_CHARACTERISTIC_CERT_LEN_UUID =
  '9B7155FC-D47E-4309-9C81-A2261D582814';

// Settings for searching
export const SERVICE_UUIDS = [SMART_LOCK_SERVICE_ID];
export const SECONDS_TO_SCAN_FOR = 7;
export const ALLOW_DUPLICATES = true;

// Other
export const CHUNK_MAX_BYTE_LENGTH = 128;
export const CHUNK_SEND_DELAY = 50; // ms
