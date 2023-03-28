#pragma once

class Console
{
private:
    static const unsigned long XTAL = 16000000;

    void sendChar(char character);
    void sendString(char *string);

    template <typename T>
    char *toString(const T &input);

public:
    Console(unsigned long baudRate, unsigned char dataBit);

    /*  template <typename T, typename... Types>
     void log(T firstArg, Types... args); */

    template <typename... Types>
    void log(const char *firstArg, Types... args);
};
