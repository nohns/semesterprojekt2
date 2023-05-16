#include "controller.h"
#include "uart.h"
#include "button.h"
#include "keypad.h"
#include "motor.h"
#include <avr/io.h>
#include <util/delay.h>

int main()
{

  // Boundary classes
  MotorDriver motor;

  // Controller classes
  Controller controller(&motor);

  // Boundary classes
  Button button(&controller);

  Uart uart();

  Keypad keypad(&controller);

  //
  while (true)
  {
    // uart.awaitRequest();
    // button.isPressed();
    // keypad.readPin();
    // x10.ProcessInput();

    motor.engageLock();
    _delay_ms(1000);

    motor.disengageLock();
    _delay_ms(1000);
  }

  return 0;
}

// https://embedds.com/controlling-servo-motor-with-avr/

/* #include <Servo.h>
#include <Arduino.h>

Servo myservo; // create servo object to control a servo
// twelve servo objects can be created on most boards

int pos = 0; // variable to store the servo position

void setup()
{
  myservo.attach(9); // attaches the servo on pin 9 to the servo object
}

void loop()
{
  for (pos = 0; pos <= 180; pos += 1)
  { // goes from 0 degrees to 180 degrees
    // in steps of 1 degree
    myservo.write(pos); // tell servo to go to position in variable 'pos'
    delay(15);          // waits 15 ms for the servo to reach the position
  }
  for (pos = 180; pos >= 0; pos -= 1)
  {                     // goes from 180 degrees to 0 degrees
    myservo.write(pos); // tell servo to go to position in variable 'pos'
    delay(15);          // waits 15 ms for the servo to reach the position
  }
} */
