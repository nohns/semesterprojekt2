#include <avr/interrupt.h>
#include <util/delay.h>

#include "controller.h"
#include "button.h"
#include "keypad.h"
#include "motor.h"
#include "x10.h"

int main()
{
  // Boundary classes
  MotorDriver motor;

  // Control classes
  Controller controller(&motor);

  // Boundary classes depending on controller
  Button button(&controller); // Button is the one that toggles the lock from the inside
  Keypad keypad(&controller);
  X10 x10;

  // Main loop
  while (true)
  {
    char x10Data = x10.readData();
    if (x10Data != X10::NO_DATA)
    {
      // The X.10 data will take the form of a 4-bit command as specified inside the routeCommand() method
      controller.routeCommand(x10Data);
    }

    // Read pin
    keypad.checkPin();
    button.checkPress();
  }

  return 0;
}
