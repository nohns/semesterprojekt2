#include <avr/io.h>
#include "x10.h"
#include <stdlib.h>
#include <avr/interrupt.h>
#include <util/delay.h>
#include <Arduino.h>

extern volatile bool zerocross;

X10::X10()
{
}

char X10 ::readData()
{

    receivedChar_ = 0;
    bitIndex_ = 4;

    while (bitIndex_ >= 0)
    {

        while (zerocross == false)
        {
            // do nothing
        }

        // If received bit is 1 on Port C pin 0
        if ((PINC & (1 << PINC0)) != 0)
        {
            receivedChar_ |= (1 << bitIndex_);
        }
        // if received bit = 0
        else
        {
            receivedChar_ &= ~(1 << bitIndex_);
        }

        bitIndex_--;
        zerocross = false;
    }

    return receivedChar_;
}

void X10 ::sendData(char c)
{

    c |= (1 << 4);
    bitIndex_ = 4;
    //Serial.println("send data");

    while (bitIndex_ >= 0)
    {

        while (zerocross == false)
        {
            // do nothing
        }
        //Serial.println("zero cross!");

        if ((c & (1 << bitIndex_)) == 0)
        {
            // set pin 6 of port H low
            PORTH &= ~(1 << PH6);
            _delay_ms(1);
            //Serial.println("low");
        }
        else
        {
            // send PWM 120kHz with 50% dc
            // initPWM();
            PORTH |= (1 << PH6);
            //_delay_us(50); //?
            _delay_ms(1);
            PORTH &= ~(1 << PH6);
            //Serial.println("high");
        }

        bitIndex_--;
        zerocross = false;
    }

    //Serial.println("-------------------");
}

void PWMDisable()
{

    OCR2A = 65000;
    OCR2B = 0;
    _delay_us(100);         // wait for OCR2A/B to change before stopping timer int
    TCCR2B &= ~(1 << CS20); // stop timer interrupt
}

void initPWM()
{
    TCCR2A = (1 << WGM21) | (1 << WGM20); // fast PWM
    TCCR2B = (1 << WGM22);                // fast PWM
    TCCR2A |= (1 << COM2B1);              // clear OC2B on Compare match
    OCR2A = 132;                          // TOP
    OCR2B = OCR2A / 2;                    // match value (50% dc)
    TCCR2B |= (1 << CS20);                // Start Timer/Counter2 with no prescaling

    _delay_ms(1);

    PWMDisable();
}