#include "controller.h"
#include "uart.h"
#include "button.h"

int main()
{

  // Boundary classes
  MotorDriver motor;

  // Controller classes
  Controller controller(&motor);

  // Boundary classes
  Button button(&controller);

  Uart uart(&controller);

  // x10Driver x10(&controller)

  while (true)
  {
    uart.awaitRequest();
    button.isPressed();
    /* x10.ProcessInput();

    */
  }

  return 0;
}