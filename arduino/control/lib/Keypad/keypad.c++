#include <avr/io.h>
#include "uart.h"
#include "keypad.h"
#include "math.h"
#include <Arduino.h>

const char READY_BYTE = 0b11001100;
const char GRANTED_BYTE = 0b00011101;
const char DENIED_BYTE = 0b11001100;

Keypad::Keypad(Controller *controller) : uart_()
{
    this->controller_ = controller;
}

void Keypad::checkPin()
{
    int pin;
    bool pinOK;

    // Waitfor pin reply from DE2-board
    this->waitForReply();

    // If pin is the same as the cached pin, then we have already verified it before, and it should not be verified again
    pin = this->readPin();
    if (pin == cachedPin_)
    {
        return;
    }
    cachedPin_ = pin;

    // Verify that the pin given is correct and send the appropriate response
    pinOK = this->controller_->verifyPin(pin);
    if (pinOK)
    {
        this->uart_->sendChar(GRANTED_BYTE);
    }
    else
    {
        this->uart_->sendChar(DENIED_BYTE);
    }
}

void Keypad::waitForReply()
{
    unsigned int counter = 0;

    // Spam the DE2-board with a little time interval with READY_BYTE until we receive a reply
    this->uart_->sendChar(READY_BYTE);
    while (!this->uart_->charReady())
    {
        if (counter < 50000)
        {
            counter++;
            continue;
        }

        counter = 0;
        this->uart_->sendChar(READY_BYTE);
    }
}

int Keypad::readPin()
{
    int pin = 0;
    int multiplier = 1; // Multiplier for digit. E.g. 1 for digit 1, 10 for digit 2 etc.

    // Read pin from DE2-board. Least significant digit first
    for (int i = 0; i < 4; i++)
    {

        // Read digit. If not in range of 0-9, then char is invalid, which means that a non 4-digit pin was entered on the DE2-board
        char digit = this->uart_->readChar();
        if (digit >= 10)
        {
            return 0;
        }

        pin += multiplier * (int)digit; // Add digit to pin
        multiplier *= 10;               // Increase multiplier for next digit
    }

    return pin;
}
