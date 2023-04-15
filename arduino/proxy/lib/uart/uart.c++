/************************************************
 * "uart.c":                                     *
 * Implementation file for Mega2560 UART driver. *
 * Using UART 0.                                 *
 * Henning Hargaard, 16/11 2022                  *
 *************************************************/
#include <avr/io.h>
#include <stdlib.h>
#include "uart.h"
#include <string.h>
#include "lock.h"
#include <stdio.h>

#define bufferLength 100

// Constructor
Uart::Uart(Controller *controller)
{
    if ((baudRate >= 300) && (baudRate <= 115200) && (dataBit >= 5) && (dataBit <= 8))
    {
        // Depency injection
        this->controller = controller;
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
    char c;
    char *buffer = (char *)malloc(sizeof(char) * bufferLength);
    while (i < bufferLength - 1)
    {
        c = readChar();
        if (c == '\x00')
            break;
        buffer[i++] = c;
    }

    //  Dynamically allocate memory for the string
    char *str = (char *)malloc(sizeof(char) * (i + 1));
    // Copy the string to the dynamically allocated memory
    strncpy(str, buffer, i + 1);
    // Free the buffer memory
    free(buffer);
    // Return the dynamically allocated memory
    return str;
}

void Uart::awaitRequest()
{
    while (true)
    {
        char *rx = readString();

        // Check if something is recieved on *rx
        if (rx != 0)
        {
            routeRequest(rx);
        }
    }
}

void Uart::routeRequest(char *request)
{

    // take a copy of the request
    char *requestCopy = (char *)malloc(strlen(request) + 1);
    strcpy(requestCopy, request);

    // Request should hold /Route/ID/State
    // Split the request into 3 parts seperated by "/"
    char *route = strtok(request, "/");
    char *id = strtok(NULL, "/");
    char *state = strtok(NULL, "/");

       if (strcmp(route, "GET") == 0)
    {
        // Create new Lock object to hold information
        Lock lock = Lock(id, state);

        // sendString("GEThehe\n");
        char *res = this->controller->getLock(lock);

        sendString(res);
        free(res);
        // sendString("GET/123/TRUE");
    }

    if (strcmp(route, "SET") == 0)
    {
        // Create new Lock object to hold information
        Lock lock = Lock(id, state);

        sendString("SEThehe\n");
        // this->controller->setLock(lock);
    }
    else
    {
        // Construct error message
        // Should consist of predefined error msg and the request route
        char *error = (char *)malloc(strlen("Error: Unrecognized request: ") + strlen(requestCopy) + 3); // +3 for "\r\n\0"
        strcpy(error, "Error: Unrecognized request: ");
        strcat(error, requestCopy);

        sendString(error);
        free(error);
        free(requestCopy);

        sprintf("Error: Unrecognized request: %s", requestCopy);
    }
    free(request);
}

/* char *request = (char *)malloc(strlen(rx) + 3); // +3 for "\r\n\0"
            strcpy(request, rx);
            strcat(request, "\r\n"); */