#pragma once
#include "uart.h"
#include "controller.h"

class Keypad
{
private:
    Controller *controller;
    Uart *uart;

    int cachedPin;

public:
    Keypad(Controller *controller);

    void readPin();
    void writeDenied();
    void writeGranted();
};