import "./style.css";

("use strict");

const bleNusServiceUUID = "6e400001-b5a3-f393-e0a9-e50e24dcca9e";
const bleNusCharRXUUID = "6e400002-b5a3-f393-e0a9-e50e24dcca9e";
const bleNusCharTXUUID = "6e400003-b5a3-f393-e0a9-e50e24dcca9e";
const MTU = 20;

let bleDevice: BluetoothDevice;
let nusService: BluetoothRemoteGATTService | undefined;
let rxCharacteristic: BluetoothRemoteGATTCharacteristic | undefined;
let txCharacteristic: BluetoothRemoteGATTCharacteristic | undefined;
let connected = false;

let fileBytes: Uint8Array;

const $connectBtn = document.getElementById(
  "clientConnectButton"
) as HTMLButtonElement;
const $uploadBtn = document.getElementById("upload") as HTMLInputElement;
const $uploadInput = document.getElementById("file") as HTMLInputElement;

// Sets button to either Connect or Disconnect
function setConnButtonState(enabled: boolean) {
  if (enabled) {
    $connectBtn.innerHTML = "Disconnect";
  } else {
    $connectBtn.innerHTML = "Connect";
  }
}

$uploadBtn.addEventListener("click", function () {
  fileBytes = new Uint8Array();
  nusSendString("u");

  setTimeout(() => {
    const blob = new Blob([fileBytes], { type: "application/octet-stream" });
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = "martin.jpg";
    a.append(document.body);
    a.target = "_blank";
    a.click();

    console.log({
      len: fileBytes.length,
    });
  }, 5000);
});

$connectBtn.addEventListener("click", function () {
  if (connected) {
    disconnect();
  } else {
    connect();
  }
});

function connect() {
  if (!navigator.bluetooth) {
    console.log(
      "WebBluetooth API is not available.\r\n" +
        "Please make sure the Web Bluetooth flag is enabled."
    );
    return;
  }

  console.log("Requesting Bluetooth Device...");
  navigator.bluetooth
    .requestDevice({
      optionalServices: [bleNusServiceUUID],
      acceptAllDevices: true,
    })
    .then((device) => {
      bleDevice = device;
      console.log("Found " + device.name);
      console.log("Connecting to GATT Server...");
      bleDevice.addEventListener("gattserverdisconnected", onDisconnected);
      return device.gatt?.connect();
    })
    .then((server) => {
      console.log("Locate NUS service");
      return server?.getPrimaryService(bleNusServiceUUID);
    })
    .then((service) => {
      nusService = service;
      console.log("Found NUS service: " + service?.uuid);
    })
    .then(() => {
      console.log("Locate RX characteristic");
      return nusService?.getCharacteristic(bleNusCharRXUUID);
    })
    .then((characteristic) => {
      rxCharacteristic = characteristic;
      console.log("Found RX characteristic");
    })
    .then(() => {
      console.log("Locate TX characteristic");
      return nusService?.getCharacteristic(bleNusCharTXUUID);
    })
    .then((characteristic) => {
      txCharacteristic = characteristic;
      console.log("Found TX characteristic");
    })
    .then(() => {
      console.log("Enable notifications");
      return txCharacteristic?.startNotifications();
    })
    .then(() => {
      console.log("Notifications started");
      txCharacteristic?.addEventListener(
        "characteristicvaluechanged",
        handleNotifications
      );
      connected = true;
      setConnButtonState(true);
    })
    .catch((error) => {
      console.log("" + error);
      if (bleDevice && bleDevice.gatt?.connected) {
        bleDevice.gatt.disconnect();
      }
    });
}

function disconnect() {
  if (!bleDevice) {
    console.log("No Bluetooth Device connected...");
    return;
  }
  console.log("Disconnecting from Bluetooth Device...");
  if (bleDevice.gatt?.connected) {
    bleDevice.gatt.disconnect();
    connected = false;
    setConnButtonState(false);
    console.log("Bluetooth Device connected: " + bleDevice.gatt.connected);
  } else {
    console.log("> Bluetooth Device is already disconnected");
  }
}

function onDisconnected() {
  connected = false;
  console.log("\r\n" + bleDevice.name + " Disconnected.");
  setConnButtonState(false);
}

function handleNotifications(event: Event) {
  console.log("notification");
  let value = (event.target as any).value as DataView;

  // Resize filebytes the amount
  const tempBuffer = new Uint8Array(fileBytes.length + value.byteLength);
  tempBuffer.set(fileBytes, 0);
  fileBytes = tempBuffer;

  const view = new DataView(
    fileBytes.buffer,
    fileBytes.byteLength - value.byteLength
  );

  console.log("value.byteLength: " + value.byteLength);
  for (let i = 0; i < value.byteLength; i++) {
    view.setUint8(i, value.getUint8(i));
  }
}

function nusSendString(s: string) {
  if (bleDevice && bleDevice.gatt?.connected) {
    console.log("send: " + s);
    let val_arr = new Uint8Array(s.length);
    for (let i = 0; i < s.length; i++) {
      let val = s[i].charCodeAt(0);
      val_arr[i] = val;
    }
    sendNextChunk(val_arr);
  } else {
    console.log("Not connected to a device yet.");
  }
}

function sendNextChunk(a: Uint8Array) {
  let chunk = a.slice(0, MTU);
  rxCharacteristic?.writeValue(chunk).then(function () {
    if (a.length > MTU) {
      sendNextChunk(a.slice(MTU));
    }
  });
}
