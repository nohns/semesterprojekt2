use bluer::{
    {adv::{Advertisement, AdvertisementHandle}, gatt::local::ApplicationHandle },
    {AdapterEvent},
    gatt::{
        local::{
            characteristic_control, Application, Characteristic, CharacteristicControlEvent,
            CharacteristicNotify, CharacteristicNotifyMethod, CharacteristicWrite, CharacteristicWriteMethod,
            Service,
        },
        CharacteristicReader, CharacteristicWriter,
    },
};

use tokio::{
    io::{ AsyncReadExt, AsyncWriteExt},
};


use futures::{future, pin_mut};
use futures_util::StreamExt;


//Cursed code



const ADVERTISER_NAME: &str = "RustenDørLås";
/// Service UUID for GATT example.
const SERVICE_UUID: uuid::Uuid = uuid::Uuid::from_u128(0xFEEDC0DE00002);

/// Characteristic UUID for GATT example.
const CHARACTERISTIC_UUID: uuid::Uuid = uuid::Uuid::from_u128(0xF00DC0DE00002);

#[tokio::main(flavor = "current_thread")]
async fn main() -> bluer::Result<()> {
    env_logger::init();
    let session = bluer::Session::new().await?;
    let adapter = session.default_adapter().await?;
    adapter.set_powered(true).await?;

//Create handle on the advertisement which can be passed along to the bluetooth struct

//Construct the bluetooth struct to hold all the data
    let bluetooth = Bluetooth::new( adapter); //This was muteable beforehand

//Start the advertisement
let handle = bluetooth.handle_advertisement().await.unwrap();

//Start the gatt server
let app_handle = bluetooth.serve_gatt_server().await;


//bluetooth.communication_handler().await;
//Start the event handler
bluetooth.event_handler(&handle).await;

    Ok(())
}

pub struct Bluetooth {
    pairing_mode: bool,
    connected: bool,
    adapter: bluer::Adapter,
}

impl Bluetooth {
   //Constructor
    pub fn new(a: bluer::Adapter ) -> Self {
       Self {
            pairing_mode: true, //Should be false by default
            connected: false,
            adapter: a,
        }
    }

    //Function to handle turning on advertisement
    async fn handle_advertisement(self: &Self) -> bluer::Result<AdvertisementHandle> {
        let advertisement = Advertisement {
            advertisement_type: bluer::adv::Type::Peripheral,
            service_uuids: vec!["123e4567-e89b-12d3-a456-426614174000".parse().unwrap()].into_iter().collect(),
            discoverable: Some(true),
            local_name: Some(ADVERTISER_NAME.to_string()),
            ..Default::default()
        };
        println!("{:?}", &advertisement);
        let handle = self.adapter.advertise(advertisement).await?;
        Ok(handle)
    }
        
//If a device connects to the device then the advertisement should be turned off
    async fn event_handler( mut self: Self, handle: &bluer::adv::AdvertisementHandle) {
        let mut events = self.adapter.events().await.unwrap();
        while let Some(event) = events.next().await {
            match event {
                AdapterEvent::DeviceAdded(adrr) => {
                    println!("Device connected: {}", adrr );
                    self.connected = true;
                    println!("connected {}" , self.connected);
                    self.pairing_mode = false;
                    println!("pairing mode {}" , self.pairing_mode);
                    //Turn off advertising
                    drop(handle)
                }
                AdapterEvent::DeviceRemoved(adrr) => {
                    println!("Device disconnected: {}", adrr);
                    self.connected = false;
                    self.pairing_mode = true;
                   //Turn on advertising again
                    self.handle_advertisement().await.unwrap();
                    println!("connected {}" , self.connected);
                    print!("pairing mode {}" , self.pairing_mode);
                    println!("handle {:?}", handle);
                }
                AdapterEvent::PropertyChanged(adapter_property) => {
                    println!("Property changed: {:?}", adapter_property);
                }
            }
        }
    }


 async fn serve_gatt_server(self: &Self) ->ApplicationHandle {
    println!("Serving GATT echo service on Bluetooth adapter {}", self.adapter.name());
    let (char_control, char_handle) = characteristic_control();
    let app = Application {
        services: vec![Service {
            uuid: SERVICE_UUID,
            primary: true,
            characteristics: vec![Characteristic {
                uuid: CHARACTERISTIC_UUID,
                write: Some(CharacteristicWrite {
                    write_without_response: true,
                    method: CharacteristicWriteMethod::Io,
                    ..Default::default()
                }),
                notify: Some(CharacteristicNotify {
                    notify: true,
                    method: CharacteristicNotifyMethod::Io,
                    ..Default::default()
                }),
                control_handle: char_handle,
                ..Default::default()
            }],
            ..Default::default()
        }],
        ..Default::default()
    };
    let app_handle = self.adapter.serve_gatt_application(app).await.unwrap();
    return app_handle;
 }

async fn communication_handler(self: &Self) {
    let mut read_buf = Vec::new();
    let mut reader_opt: Option<CharacteristicReader> = None;
    let mut writer_opt: Option<CharacteristicWriter> = None;
    let (char_control, char_handle) = characteristic_control();
    pin_mut!(char_control);
    loop {
        tokio::select! {
            evt = char_control.next() => {
                match evt {
                    Some(CharacteristicControlEvent::Write(req)) => {
                        println!("Accepting write request event with MTU {}", req.mtu());
                        read_buf = vec![0; req.mtu()];
                        reader_opt = Some(req.accept().unwrap());
                    },
                    Some(CharacteristicControlEvent::Notify(notifier)) => {
                        println!("Accepting notify request event with MTU {}", notifier.mtu());
                        writer_opt = Some(notifier);
                    },
                    None => break,
                }
            },
            read_res = async {
                match &mut reader_opt {
                    Some(reader) if writer_opt.is_some() => reader.read(&mut read_buf).await,
                    _ => future::pending().await,
                }
            } => {
                match read_res {
                    Ok(0) => {
                        println!("Read stream ended");
                        reader_opt = None;
                    }
                    Ok(n) => {
                        let value = read_buf[..n].to_vec();
                        println!("Echoing {} bytes: {:x?} ... {:x?}", value.len(), &value[0..4.min(value.len())], &value[value.len().saturating_sub(4) ..]);
                        if value.len() < 512 {
                            println!();
                        }
                        if let Err(err) = writer_opt.as_mut().unwrap().write_all(&value).await {
                            println!("Write failed: {}", &err);
                            writer_opt = None;
                        }
                    }
                    Err(err) => {
                        println!("Read stream error: {}", &err);
                        reader_opt = None;
                    }
                }
            }
        }
    }
}
   




}

//scp tooth martin@raspberrypi:~/

