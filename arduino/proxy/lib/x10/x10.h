#pragma once

class X10
{
private:
    char receivedChar_;
    int bitIndex_;

public:
    X10();
    char readData();
    void sendData(char);
    void PWMDisable();
    void initPWM();
};
