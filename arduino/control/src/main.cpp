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

<<<<<<< HEAD
  Keypad keypad (&controller);

//
=======
  // x10Driver x10(&controller)

>>>>>>> 07bdea21c318e6fe83c54c54909d1f8c0582500b
  while (true)
  {
    //uart.awaitRequest();
    button.isPressed();
    keypad.readPin();
    /* x10.ProcessInput();
<<<<<<< HEAD
    
=======

>>>>>>> 07bdea21c318e6fe83c54c54909d1f8c0582500b
    */
  }

  return 0;
}