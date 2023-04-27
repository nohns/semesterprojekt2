#include "controller.h"
#include "uart.h"

int main()
{

  /*   MotorDriver motor;

    Control control(&motor);

    x10Driver x10(&control);

    UartDriver uart(&control);

    Button button(&control); */

  Controller controller;

  Uart uart(&controller);

  for (;;)
  {
    uart.awaitRequest();
    /* x10.ProcessInput();
    uart.ProcessInput();
    button.ProcessInput(); */
  }

  return 0;
}