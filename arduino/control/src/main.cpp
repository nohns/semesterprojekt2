#include "controller.h"
#include "button.h"
#include "keypad.h"
#include "motor.h"
#include "uart.h"

#include <util/delay.h>
// #include "x10.h"

int main()
{

  // Boundary classes
  MotorDriver motor;

  // Controller classes
  Controller controller(&motor);

  // Boundary classes
  Button button(&controller);

  Keypad keypad(&controller);

  // X10
  // X10 x10(&controller);

  Uart uart;

  //
  while (true)
  {
    button.isPressed();
    keypad.readPin();
    _delay_ms(10);
  }

  return 0;
}