
#include "uart.h"
#include "controller.h"
#include "x10.h"

#include <stdio.h>
#include <stdlib.h>

int main()
{

  X10 x10;

  Controller controller(&x10);

  Uart uart(&controller);

  // Start uart eventHandler
  while (true)
  {
    uart.awaitRequest();
  }
  return 0;
}