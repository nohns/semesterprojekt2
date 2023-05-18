#pragma once

#include "controller.h"

class Uart
{
private:
    // Constants to make sure its compatible with Go and C++
    // They have to be unsigned longs because of mafs (maths) reasons
    static const unsigned long XTAL = 16000000;
    static const unsigned long baudRate = 9600;
    static const unsigned char dataBit = 8;

    Controller *controller;

    // Internal use functions
    bool charReady();
    void sendChar(char character);
    char readChar();

public:
    Uart(Controller *controller);
    Uart();

    void awaitRequest();
    void sendCommand(char cmd);
    char readCommand();

    // Used for testing only
    void sendString(char *string);
};
