

#include <avr/io.h>

#include "button.h"
#include "controller.h"

Button::Button(Controller *controller)
{
    this->controller_ = controller;
    this->wasPressed_ = false;

    // Set pin 7 to input with pull-up
    DDRA &= ~(1 << PINA7);
    PORTA |= (1 << PINA7);
}

void Button::checkPress()
{
    // If pressed now is different from pressed last time, then we have a new button press
    bool isPressedNow = (PINA & (1 << PINA7)) == 0;
    if (isPressedNow != wasPressed_)
    {
        // Only toggle when button just got pressed (falling edge, as button is active low)
        if (isPressedNow)
        {
            this->controller_->toggleLock();
        }
    }

    wasPressed_ = isPressedNow;
}