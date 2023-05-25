#pragma once

#include "controller.h"

class Button
{
private:
    Controller *controller_;
    bool wasPressed_;

public:
    Button(Controller *controller);

    // Check if the button is pressed and toggle the lock if it is
    void checkPress();
};