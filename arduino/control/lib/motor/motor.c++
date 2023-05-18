#include "motor.h"

#include <util/delay.h>
#include <avr/io.h>

// Constructor to create motor object
/* MotorDriver::MotorDriver()
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
    _delay_ms(50);
    OCR1A = 0;
}

void MotorDriver::disengageLock()
{
    OCR1A = 39999 - 1999;

    _delay_ms(50);
    OCR1A = 0;
} */

// TEST: CODE
MotorDriver::MotorDriver()
{
    // Set up I/O pin for output
    DDRB = 0xff; // Output pin is set as pin 11

    TCCR1A |= (1 << COM1A1) | (1 << COM1B1) | (1 << WGM11);
    TCCR1B |= (1 << WGM13) | (1 << WGM12) | (1 << CS11) | (1 << CS10);

    ICR1 = 4999; // Set ICR1 for 20ms signal period
}

void MotorDriver::engageLock()
{
    OCR1A = 535;
}

void MotorDriver::disengageLock()
{
    OCR1A = 97;
}
