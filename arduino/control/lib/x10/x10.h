#include <Arduino.h>

#pragma once

class X10
{
private:
    char receivedChar_;

public:
    X10();
    char readData();
    void sendData(char);
};
