#include "controller.h"
#include "button.h"
#include "keypad.h"
#include "motor.h"
#include "x10.h"
#include <avr/interrupt.h>
#include <util/delay.h>

volatile bool zerocross;
volatile int bitIndex = 4;

void intExInterrupt();

int main()
{

  DDRC = 0x00;

  intExInterrupt();

  // Boundary classes
  MotorDriver motor;

  // Controller classes
  Controller controller(&motor);

  // Boundary classes
  Button button(&controller);

  Keypad keypad(&controller);

  // X10
  X10 x10;

  //
  while (true)
  {
    char output = x10.readData();
    controller.routeCommand(output);
    keypad.readPin();

    button.isPressed();
  }

  return 0;
}

void intExInterrupt()
{
  EIMSK |= 0b00000001;
  EICRA |= 0b00000011;
  sei();
}

ISR(INT0_vect)
{
  zerocross = true;
}
