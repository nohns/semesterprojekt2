#include <avr/io.h>
#include <stdlib.h>
#include <avr/interrupt.h>
#include <util/delay.h>

#include "x10.h"

volatile bool zeroCross = false;

X10::X10()
{
    DDRC &= ~(1 << PINC0); // Set pin 0 on port C to input
    DDRH |= (1 << PH6);    // Set pin 6 on port H to output

    // Init external interrupt INT0 on rising edge for x10 zero cross signal
    EIMSK |= 0b00000001;
    EICRA |= 0b00000011;
    sei();

    // Enable 120kHz PWM on timer 2
    TCCR2A = (1 << WGM21) | (1 << WGM20); // fast PWM
    TCCR2B = (1 << WGM22);                // fast PWM
    TCCR2A |= (1 << COM2B1);              // clear OC2B on Compare match
}

char X10::readData()
{
    int readBitIndex = 4;
    char recvChar = NO_DATA;

    // Loop until we have read all 4 bits (5 with startbit) from X.10
    while (readBitIndex >= 0)
    {
        // Wait for a zero cross
        while (zeroCross == false)
        {
        }
        // Delay to make sure the sender actually have had time to send the bit
        _delay_us(100);

        // Break out of read loop, if no start bit has been sent (index 4, which means byte number 5 in the sequence)
        // This will make sure that we just check the start bit again at next zero cross, but at the same time allow all other
        // logic to continue running, e.g. checking for button presses, pincode etc.
        if (readBitIndex == 4 && (PINC & (1 << PINC0)) == 0)
        {
            zeroCross = false;
            break;
        }

        // If received bit is 0 on Port C pin 0
        if ((PINC & (1 << PINC0)) != 0)
        {
            recvChar |= (1 << readBitIndex);
        }
        // if received bit != 0
        else
        {
            recvChar &= ~(1 << readBitIndex);
        }

        // Make loop ready for recording the next bit and for waiting until next zero cross
        readBitIndex--;
        zeroCross = false;
    }

    // Set startbit to 0 in receivedChar
    recvChar &= ~(1 << 4);
    return recvChar;
}

void X10::sendData(char c)
{
    int sendBitIndex = 4;

    // Set startbit at bit number 5 (index 4)
    c |= (1 << 4);

    // Loop until we have sent all 4 bits (5 with startbit) over X.10
    while (sendBitIndex >= 0)
    {
        // Wait for a zero cross to occur before sending anything
        while (zeroCross == false)
        {
        }

        // If bit to send is 0, make sure pin 6 of port H is low
        if ((c & (1 << sendBitIndex)) == 0)
        {
            // Set pin 6 of port H low
            PORTH &= ~(1 << PH6);
            _delay_ms(1);
        }
        else
        {
            // Send PWM 120kHz with 50% dc
            this->startPWM();
            // PORTH |= (1 << PH6); // USED FOR SIMULATING SIGNAL OVER ANALOG DISCOVERY
            _delay_ms(1);
            this->endPWM();
            // PORTH &= ~(1 << PH6); // USED FOR SIMULATING SIGNAL OVER ANALOG DISCOVERY
        }

        sendBitIndex--;
        zeroCross = false;
    }
}

void X10::startPWM()
{
    // Set duty cycle to 50%
    OCR2A = 132;       // TOP
    OCR2B = OCR2A / 2; // match value (50% dc)

    // Start timer 2 with no prescaling
    TCCR2B |= (1 << CS20); // Start Timer/Counter2 with no prescaling
}

void X10::endPWM()
{
    // Set duty cycle to 0
    OCR2A = 65000;
    OCR2B = 0;

    // Wait for OCR2A/B to change before disabling timer
    _delay_us(100);

    // Sisable timer 2
    TCCR2B &= ~(1 << CS20);
}

ISR(INT0_vect)
{
    zeroCross = true;
}