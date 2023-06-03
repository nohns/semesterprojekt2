#pragma once

class X10
{
private:
 // Start 120kHz PWM signal for outputting a bit as a 1
    void startPWM();
    // End PWM signal
    void endPWM();

public:
    X10();
    char readData();
    void sendData(char);
   
       // Byte code that represent no data received from X.10
    static const char NO_DATA = 0xFF;
};
