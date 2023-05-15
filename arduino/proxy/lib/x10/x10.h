#pragma once

class X10
{
private:
    bool dataHigh_;
    char receivedChar_;

public:
    X10();
    char readData();
    void sendData(char);
    void setDataHigh(bool);
    void setReceivedCharHigh(int);
    void setReceivedCharLow(int);
};
