

#include <avr/io.h>

#include "Button.h"
#include "controller.h"

// Constructor
Button::Button(Controller *controller)
{
    // Dependency inject the controller
    this->controller = controller;

    // set pin 7 to input
    DDRA = 0;
}

// Method to check if the hardware button is pressed
// returns true if button is pressed and false if not
bool Button::isPressed()
{

    if ((PINA & 0b10000000) == 0)
    {
        return true;
    };

    return false;
}