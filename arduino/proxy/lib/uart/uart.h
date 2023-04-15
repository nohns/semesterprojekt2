#pragma once

#include "controller.h"

class Uart
{
private:
    // Constants to make sure its compatible with Go and C++
    static const unsigned long XTAL = 16000000;
    static const unsigned long baudRate = 9600;
    static const unsigned char dataBit = 8;

    Controller *controller;

    // Both of these functions are for internal use only
    bool
    charReady();
    void sendChar(char character);
    char readChar();

public:
    Uart(Controller *controller);

    char *readString();
    void sendString(char *string);

    void awaitRequest();
    void routeRequest(char *request);
};
