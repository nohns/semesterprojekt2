#include "controller.h"
#include "uart.h"
#include "button.h"
#include "keypad.h"

int main()
{

  // Boundary classes
  MotorDriver motor;

  // Controller classes
  Controller controller(&motor);

  // Boundary classes
  Button button(&controller);

  Uart uart();

  Keypad keypad (&controller);

//
  while (true)
  {
    //uart.awaitRequest();
    button.isPressed();
    keypad.readPin();
    /* x10.ProcessInput();
    
    */
  }

  return 0;
}