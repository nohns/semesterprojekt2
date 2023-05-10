#include "motor.h"

#include <util/delay.h>

// Constructor to create motor object
MotorDriver::MotorDriver()
{
    // Set up I/O pin for output
    DDRB = 0xff; // Output pin is set as pin 11

    TCCR1A = 0b11000010;
    TCCR1B = 0b00011001;

    ICR1 = 39999; // Set ICR1 for 20ms signal period
}

void MotorDriver::engageLock()
{
    OCR1A = 39999 - 3999;
    _delay_ms(100);
}

void MotorDriver::disengageLock()
{
    OCR1A = 39999 - 1999;

    _delay_ms(100);
}