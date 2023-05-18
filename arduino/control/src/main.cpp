#include "controller.h"
#include "button.h"
#include "keypad.h"
#include "motor.h"
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

  //
  while (true)
  {
    button.isPressed();
    keypad.readPin();
    // x10.ProcessInput();
  }

  return 0;
}