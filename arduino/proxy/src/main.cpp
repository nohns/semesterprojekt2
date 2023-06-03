#include "x10.h"
#include "controller.h"
#include "uart.h"
#include <stdio.h>
#include <stdlib.h>
#include <avr/interrupt.h>
#include <util/delay.h>


int main(void)
{

  X10 x10;

  Controller Controller(&x10);

  Uart uart(&Controller);

  while (true)

  {
    uart.awaitRequest();
  }

  return 0;
}