#include <avr/io.h>
#include <stdlib.h>
#include "uart.h"

#include <string.h>
#include <stdio.h>

#define bufferLength 100

// Constructor
Uart::Uart(Controller *controller)
{
    if ((baudRate >= 300) && (baudRate <= 115200) && (dataBit >= 5) && (dataBit <= 8))
    {
        UCSR0B = 0b00011000;                // Asynchronous operation, 1 stop bit
        UCSR0C = (dataBit - 5) << 1;        // Bit 2 and bit 1 controls the number of data bits
        UCSR0C &= ~(1 << UPM01);            // Set parity to None
        UCSR0C &= ~(1 << UPM00);            // Set parity to None
        UBRR0 = XTAL / (16 * baudRate) - 1; // Set Baud Rate according to the parameter baudRate

        // Depency injection
        this->controller = controller;
    }
}

// Constructor without dependency injection - used for testing
Uart::Uart()
{
    if ((baudRate >= 300) && (baudRate <= 115200) && (dataBit >= 5) && (dataBit <= 8))
    {
        UCSR0B = 0b00011000;                // Asynchronous operation, 1 stop bit
        UCSR0C = (dataBit - 5) << 1;        // Bit 2 and bit 1 controls the number of data bits
        UCSR0C &= ~(1 << UPM01);            // Set parity to None
        UCSR0C &= ~(1 << UPM00);            // Set parity to None
        UBRR0 = XTAL / (16 * baudRate) - 1; // Set Baud Rate according to the parameter baudRate
    }
}

// Check if UART has received a new character
bool Uart::charReady()
{
    return UCSR0A & (1 << 7);
}

// Send a character over UART
void Uart::sendChar(char character)
{
    // Wait for transmitter register empty (ready for new character)
    while ((UCSR0A & (1 << 5)) == 0)
    {
    }
    // Then send the character
    UDR0 = character;
}

// Used for debugging mainly
void Uart::sendString(char *str)
{
    // Repeat until zero-termination
    while (*str != 0)
    {
        // Send the character pointed to by "string"
        sendChar(*str);
        // Advance the pointer one step
        str++;
    }
    // Send null terminating byte to indicate end of string
    sendChar('\x00');
}

void Uart::sendCommand(char cmd)
{
    // The command is contained within a single character
    sendChar(cmd);
    // Send null terminating byte to indicate end of command
    sendChar('\x00');
}

// Read a character from UART
char Uart::readChar()
{
    // Wait for new character received
    while ((UCSR0A & (1 << 7)) == 0)
    {
    }
    // Then return it
    return UDR0;
}

// Remember to free the memory after use
char Uart::readCommand()
{
    int i = 0;
    char buffer[bufferLength];
    while (i < bufferLength - 1)
    {
        char c = readChar();
        if (c == '\x00')
        {
            break;
        }
        buffer[i++] = c;
    }
    return buffer[0];
}

void Uart::awaitRequest()
{
    char rx = readCommand();
    // Check if something is recieved on *rx
    if (rx != 0)
    {
        // Call controller to route request
        char res = this->controller->forwardRequest(rx);

        // Send response back to bridge
        sendCommand(res);

        // Release the malloced memory
    }
}
