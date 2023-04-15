
#include "uart.h"
#include "controller.h"
#include <stdio.h>
#include <stdlib.h>

int main()
{
  Controller controller;

  Uart uart(&controller);

  // Start uart evenHandler
  uart.awaitRequest();

  return 0;
}