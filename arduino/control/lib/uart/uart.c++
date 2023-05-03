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

        // No interrupts enabled
        // Receiver enabled
        // Transmitter enabled
        // No 9 bit operation
        // UCSR0B = 0b00011000;
        // Asynchronous operation, 1 stop bit
        // Bit 2 and bit 1 controls the number of data bits
        // UCSR0C = (dataBit - 5) << 1;
        // Set Baud Rate according to the parameter baudRate
        // UBRR0 = XTAL / (16 * baudRate) - 1;

        UCSR0B = 0b00011000;                // Asynchronous operation, 1 stop bit
        UCSR0C = (dataBit - 5) << 1;        // Bit 2 and bit 1 controls the number of data bits
        UCSR0C &= ~(1 << UPM01);            // Set parity to None
        UCSR0C &= ~(1 << UPM00);            // Set parity to None
        UBRR0 = XTAL / (16 * baudRate) - 1; // Set Baud Rate according to the parameter baudRate

        // Depency injection
        this->controller = controller;
    }
}

/*  UCSR0B = 0b00011000;
    // Enable RX interrupt (if required by parameter)
    if (RX_Int)
      UCSR0B |= (1<<7);
    // Asynchronous operation, 1 stop bit
    // Bit 2 and bit 1 controls the number of data bits
    UCSR0C = (DataBit-5)<<1;
    // Set Baud Rate according to the parameter BaudRate
    UBRR0 = XTAL/(16*BaudRate) - 1; */

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

// Send a null-terminated string over UART
void Uart::sendString(char *string)
{
    // Repeat until zero-termination
    while (*string != 0)
    {
        // Send the character pointed to by "string"
        sendChar(*string);
        // Advance the pointer one step
        string++;
    }
    sendChar('\x00');
}

void Uart::sendCommand(char cmd)
{
    sendChar(cmd);
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
char *Uart::readString()
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

    //  Dynamically allocate memory for the string
    char *str = (char *)malloc(sizeof(char));
    if (str != NULL)
    {
        //  Copy the string to the dynamically allocated memory
        strncpy(str, buffer, sizeof(char));
        return str;
    }

    return NULL;
}

void Uart::awaitRequest()
{
    char *rx = readString();
    // Check if something is recieved on *rx
    if (rx != 0)
    {
        // Call controller to route request
        // char *res = this->controller->forwardRequest(rx);

        // Send response back to bridge
        // sendCommand(res[0]);

        // free(res);
        free(rx);
    }
}
