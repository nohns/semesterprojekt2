#include "controller.h"
#include "uart.h"
#include "button.h"
#include "keypad.h"

int main()
{

  /*
      Control control(&motor);

      x10Driver x10(&control);

      UartDriver uart(&control);

      Button button(&control); */

  MotorDriver motor;

  Controller controller(&motor);

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