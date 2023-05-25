#pragma once

#include "uart.h"
#include "controller.h"

class Keypad
{
private:
    Controller *controller_;
    Uart *uart_;

    // Cached pin to make sure we only verify once per pin read
    int cachedPin_;

    // Waits until we receive a reply from the DE2-board
    void waitForReply();

    // Reads current pin as an integer from the DE2-board
    int readPin();

public:
    Keypad(Controller *controller);

    // Perform a pin check with the current pin entered on the DE2-board against the pin stored in the domain
    void checkPin();
};