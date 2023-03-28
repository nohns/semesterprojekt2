
#include "uart.h"
#include "logging.h"

int main()
{
  Uart uart(9600, 8);

  // Create an event handler which listens for events on uart.ReadChar()

  for (;;)
  {
    uart.sendString("Hello World!");
    char rx = uart.readChar();
    if (rx == 'a')
    {
      uart.sendString("You pressed a");
    }
    else if (rx == 'b')
    {
      uart.sendString("You pressed b");
    }
    else if (rx == 'c')
    {
      uart.sendString("You pressed c");
    }
  }

  return 0;
}