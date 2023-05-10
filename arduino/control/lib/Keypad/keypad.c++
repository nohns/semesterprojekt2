#include <avr/io.h>
#include "uart.h"
#include "keypad.h"
#include "math.h"

#define READY_BYTE 0b11001100
#define Granted_BYTE 0b00011101
#define Denied_BYTE 0b11001100

Keypad::Keypad(Controller *controller)
{
    Uart uart;
    this->uart = &uart;
}

int Keypad::readPin()
{
    // Sending ready bits
    this->uart->sendChar(READY_BYTE);

    int pin = 0;
    for (int i = 0; i < 4; i++)
    {
        // Læs pinchar. Kan være 1-10 hvor 10 betyder at det i'te cifret ikke er givet
        char digit = this->uart->readChar();
        if (digit >= 10)
        {
            continue;
        }

        // pin += 10^i * pinchar
        pin += pow(10, i) * digit;
    }

    return pin;
}

void Keypad::writeDenied()
{

    this->uart->sendChar(Denied_BYTE);
}

void Keypad::writeGranted()
{

    this->uart->sendChar(Granted_BYTE);
}
