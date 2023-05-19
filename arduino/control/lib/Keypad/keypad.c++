#include <avr/io.h>
#include "uart.h"
#include "keypad.h"
#include "math.h"
#include <Arduino.h>

#define READY_BYTE 0b11001100
#define Granted_BYTE 0b00011101
#define Denied_BYTE 0b11001100

Keypad::Keypad(Controller *controller) : uart()
{
    this->controller = controller;
}

void Keypad::readPin()
{

    this->uart->sendChar(READY_BYTE);

    // CURSED WORK AROUND for getting the uart to respond in time
    long counter = 0;
    while (!this->uart->charReady())
    {
        if (counter < 50000)
        {
            counter++;
            continue;
        }

        counter = 0;
        this->uart->sendChar(READY_BYTE);
    }

    char digit1 = this->uart->readChar();
    char digit2 = this->uart->readChar();
    char digit3 = this->uart->readChar();
    char digit4 = this->uart->readChar();
    if (digit1 >= 10 || digit2 >= 10 || digit3 >= 10 || digit4 >= 10)
    {
        return;
    }

    int pin = digit1 + digit2 * 10 + digit3 * 100 + digit4 * 1000;

    Serial.print("pin recv: ");
    Serial.println(pin);

    // Just send it to make de2 happy
    if (pin == cachedPin)
    {
        return;
    }
    cachedPin = pin;
    Serial.println("pin changed");

    bool ok = this->controller->verifyPin(pin);
    /*if (ok)
    {
        this->writeGranted();
    }
    if (!ok)
    {
        this->writeDenied();
    }*/
}

void Keypad::writeDenied()
{

    this->uart->sendChar(Denied_BYTE);
}

void Keypad::writeGranted()
{

    this->uart->sendChar(Granted_BYTE);
}
