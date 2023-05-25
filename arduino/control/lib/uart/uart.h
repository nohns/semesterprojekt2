#pragma once

class Uart
{
private:
    // Constants to make sure its compatible with Go and C++
    // They have to be unsigned long else shit breaks dont ask me why
    static const unsigned long XTAL = 16000000;
    static const unsigned long baudRate = 9600;
    static const unsigned char dataBit = 8;

public:
    Uart();

    bool charReady();
    void sendChar(char character);
    char readChar();
};
