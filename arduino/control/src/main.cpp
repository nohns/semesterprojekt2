#include "controller.h"
#include "uart.h"
#include "button.h"

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

  Uart uart(&controller);

  while (true)
  {
    uart.awaitRequest();
    button.isPressed();
    /* x10.ProcessInput();
    uart.ProcessInput();
    */
  }

  return 0;
}