
#include "uart.h"
#include "logging.h"

int main()
{
  Uart uart(9600, 8);

  // Create an event handler which listens for events on uart.ReadChar()

  for (;;)
  {

    // char rx = uart.readChar();

    char buffer[100];
    uart.readString(buffer, sizeof(buffer));
    uart.sendString(buffer);

    if (buffer == "123")
    {
      uart.sendString("123");
    }
    else if (buffer == "b")
    {
      uart.sendString("You pressed b");
    }
  }

  return 0;
}