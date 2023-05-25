#include <util/delay.h>
#include <avr/io.h>

#include "motor.h"

MotorDriver::MotorDriver()
{
    // Set up I/O pin for output
    DDRB = (1 << PB5); // Output pin is set as pin 11

    // Set up timer 1 for PWM
    TCCR1A |= (1 << COM1A1) | (1 << COM1B1) | (1 << WGM11);
    TCCR1B |= (1 << WGM13) | (1 << WGM12) | (1 << CS11) | (1 << CS10);
    ICR1 = ICR_20MS_PERIOD; // Set ICR1 for 20ms signal period
}

void MotorDriver::engageLock()
{
    OCR1A = OCR_LOCKED;
}

void MotorDriver::disengageLock()
{
    OCR1A = OCR_UNLOCKED;
}
