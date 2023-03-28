/************************************************
 * "uart.c":                                     *
 * Implementation file for Mega2560 UART driver. *
 * Using UART 0.                                 *
 * Henning Hargaard, 16/11 2022                  *
 *************************************************/
#include <avr/io.h>
#include <stdlib.h>
#include "uart.h"

// Constructor
Uart::Uart(unsigned long baudRate, unsigned char dataBit)
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
    }
}

// Check if UART has received a new character
bool Uart::charReady()
{
    return UCSR0A & (1 << 7);
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
}

// Convert an integer to a string and send it over UART
void Uart::sendInteger(int number)
{
    char array[7];
    // Convert the integer to an ASCII string (array), radix = 10
    itoa(number, array, 10);
    // - then send the string
    sendString(array);
}
