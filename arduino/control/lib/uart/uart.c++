#include <avr/io.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <avr/interrupt.h>

#include "uart.h"

#define bufferLength 100

// Constructor
Uart::Uart()
{
    if ((baudRate >= 300) && (baudRate <= 115200) && (dataBit >= 5) && (dataBit <= 8))
    {
        UCSR1B = 0b00011000;                // Enable RX and TX
        UCSR1C = (dataBit - 5) << 1;        // Bit 2 and bit 1 controls the number of data bits
        UCSR1C &= ~(1 << UPM11);            // Set parity to None
        UCSR1C &= ~(1 << UPM10);            // Set parity to None
        UBRR1 = XTAL / (16 * baudRate) - 1; // Set Baud Rate according to the parameter baudRate
    }
}

// Check if UART has received a new character
bool Uart::charReady()
{
    return UCSR1A & (1 << 7);
}

// Send a character over UART
void Uart::sendChar(char character)
{
    // Wait for transmitter register empty (ready for new character)
    while ((UCSR1A & (1 << 5)) == 0)
    {
    }
    // Then send the character
    UDR1 = character;
}

// Read a character from UART
char Uart::readChar()
{
    // Wait for new character received
    while ((UCSR1A & (1 << 7)) == 0)
    {
    }
    // Then return it
    return UDR1;
}
