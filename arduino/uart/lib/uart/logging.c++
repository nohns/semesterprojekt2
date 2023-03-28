#include "logging.h"
#include <avr/io.h>
#include <stdlib.h>

Console::Console(unsigned long baudRate, unsigned char dataBit)
{
    if ((baudRate >= 300) && (baudRate <= 115200) && (dataBit >= 5) && (dataBit <= 8))
    {
        // No interrupts enabled
        // Receiver enabled
        // Transmitter enabled
        // No 9 bit operation
        UCSR0B = 0b00011000;
        // Asynchronous operation, 1 stop bit
        // Bit 2 and bit 1 controls the number of data bits
        UCSR0C = (dataBit - 5) << 1;
        // Set Baud Rate according to the parameter baudRate
        UBRR0 = XTAL / (16 * baudRate) - 1;
    }
}

// Send a character over UART
void Console::sendChar(char character)
{
    // Wait for transmitter register empty (ready for new character)
    while ((UCSR0A & (1 << 5)) == 0)
    {
    }
    // Then send the character
    UDR0 = character;
}

// Send a null-terminated string over UART
void Console::sendString(char *input)
{
    if (input == nullptr)
    {
        return;
    }
    // Repeat until zero-termination
    while (*input != 0)
    {
        // Send the character pointed to by "string"
        sendChar(*input);
        // Advance the pointer one step
        input++;
    }
}

// This function converts the argument message to a string.
// The argument message can be of any type.
// The function returns a string.

template <typename T>
char *Console::toString(const T &input)
{
}

// sendString converts the argument to a string and sends it to the console
/* template <typename T, typename... Types>
void Console::log(T firstArg, Types... args)
{
    sendString(toString(firstArg));
    log(args...);
} */

template <typename... Types>
void Console::log(const char *firstArg, Types... args)
{
    sendString(toString(firstArg));
    log(args...);
}