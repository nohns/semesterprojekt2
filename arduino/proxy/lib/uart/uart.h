#pragma once

class Uart
{
private:
    // Target CPU frequency
    static const unsigned long XTAL = 16000000;

    // Both of these functions are for internal use only
    bool charReady();
    void sendChar(char character);

public:
    Uart(unsigned long baudRate, unsigned char dataBit);

    void readString(char *buffer, int bufferLength);

    char readChar();
    void sendString(char *string);
    void sendInteger(int number);
};
