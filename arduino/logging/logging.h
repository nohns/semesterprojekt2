#pragma once

#include <string>

class Console
{
private:
    void sendChar(char character);
    void sendString(char *string);

    template <typename T>
    std::string Console::toString(const T &input);

public:
    Console(unsigned long baudRate, unsigned char dataBit);

    template <typename T, typename... Types>
    void log(T firstArg, Types... args);
};
