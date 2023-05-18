

#include <avr/io.h>

#include "button.h"
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
void Button::isPressed()
{

    if ((PINA & 0b10000000) == 0)
    {
        // If we detect an input we should call the controller and request a toggle on the lock
        controller->toggleLock();
    };
}