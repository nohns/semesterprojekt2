#include <avr/io.h>
#include "x10.h"
#include <stdlib.h>
#include <avr/interrupt.h>

extern volatile bool zerocross;
extern volatile bool flag;
extern volatile int bitIndex;

X10::X10()
{
    receivedChar_ = 0;
    dataHigh_ = false;
}

void X10 ::setDataHigh(bool datahigh)
{
    dataHigh_ = datahigh;
}

void X10 ::setReceivedCharHigh(int bitindex)
{
    receivedChar_ |= (1 << bitindex);
}

void X10 ::setReceivedCharLow(int bitindex)
{
    receivedChar_ |= ~(1 << bitindex);
}

char X10 ::readData()
{
    /* do
    { */
    while (zerocross == false)
    {
        // do nothing
    }
    // set Timer 2 to generate interrupts after period for 120 kHz
    /* TCCR2A = 0;             // set Timer 2 to normal mode
    TCCR2B = (1 << CS20);   // set prescaler to 1
    OCR2A = 83;             // set TOP value to 83 for 120 kHz frequency
    TIMSK2 = (1 << OCIE2A); // enable Timer 2 compare match interrupt
*/
    /* while (flag == false)
    { */

    if (bitIndex >= 0)
    {
        // If received bit is 0 on Port C pin 0
        if ((bitIndex >= 0) && ((PINC & (0 << PINC0)) == (0 << PINC0)))
        {
            setReceivedCharLow(bitIndex);
            bitIndex--;
            zerocross = false;
        }
        // if received bit != 0
        else
        {
            setReceivedCharHigh(bitIndex);
            bitIndex--;
            zerocross = false;
            /* setDataHigh(true); */
        }
    }
    else
    {
        bitIndex = 4;
        zerocross = false;
        return 1;
    }

    /*  } */

    /* flag = false; */

    // set char bit value to received data
    /* if (dataHigh_ == true)
    { */
    /* setReceivedCharHigh(bitIndex); */
    /* setDataHigh(false); */
    /* } */
    /* else */
    /* { */
    /* setReceivedCharLow(bitIndex); */
    /*  } */

    /* } while (bitIndex >= 0); */

    /* bitIndex = 4; */

    // return receivedChar if startbit is received
    return /* (receivedChar_ & (1 << 4) ? receivedChar_ : 0) */ receivedChar_;
}

void X10 ::sendData(char c)
{
    do
    {
        while (zerocross == false)
        {
            // do nothing just wait
        }

        // Set Timer 1 to CTC mode and enable interrupt
        TCCR1A = 0;
        TCCR1B = (1 << WGM12) | (1 << CS12) | (1 << CS10);
        OCR1A = 15;
        TIMSK1 = (1 << OCIE1A);

        while (flag == false)
        {
            // if start bit is one
            if ((c & (1 << 4)))
            {
                // do nothing
            }

            else
            { // set startbit to one
                c |= (1 << 4);
            }

            if ((c & (1 << bitIndex)) == 0)
            {
                // set pin 6 of port H low
                PORTH &= ~(1 << PH6);
            }
            else
            {
                // sets pin 6 of port H to 120kHz
                // PWM at 120 kHz with DC=50%
                // Timer 2 PWM fast mode enables and non-inverting mode is set and prescaler 8
                TCCR2A = 0b10000011;
                TCCR2B = 0b00001010;
                // TOP LEVEL set to 128 Timer 2
                OCR2A = 127;
                // OCR value is set to 64 for Timer 2
                OCR2B = 64;
            }
        }
        flag = false;

    } while (bitIndex >= 0);

    bitIndex = 4;
}

// Define timer ISR to be executed every 1 ms
ISR(TIMER1_COMPA_vect)
{
    flag = true;
    bitIndex--;
    zerocross = false;
    // disable PWM
    TCCR2A = 0;
    TCCR2B = 0;
    // disable timer compare interrupt
    TIMSK1 &= ~(1 << OCIE1A);
}

ISR(TIMER2_COMPA_vect)
{
    flag = true;
    bitIndex--;
    zerocross = false;
    TCCR2B = 0; // disable Timer 2
}