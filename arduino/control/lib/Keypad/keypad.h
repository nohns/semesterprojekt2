#pragma once
#include "uart.h"
#include "controller.h"

class Keypad
{
private:
    Controller *controller;
    Uart * uart;
public:
Keypad(Controller *controller );

    int readPin();
    void writeDenied();
    void writeGranted();
};