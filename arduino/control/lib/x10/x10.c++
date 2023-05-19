#include <avr/io.h>
#include "x10.h"
#include <stdlib.h>
#include <avr/interrupt.h>
#include <util/delay.h>

extern volatile bool zerocross;
extern volatile int bitIndex;

// create global serial object

X10::X10()
{
}

char X10 ::readData()
{
    int readBitIndex = 4;
    char recvChar = 0;

    while (readBitIndex >= 0)
    {
        // Serial.println("Bit index: ");
        // Serial.println(bitIndex);
        while (zerocross == false)
        {
            // do nothing
            // Serial.println("ZeroCross false");
        }
        // Delay to make sure the sender actually have time to send the bit
        _delay_us(100);

        // Wait for a start bit (index 4, which means byte number 5)
        while (readBitIndex == 4 && (PINC & (1 << PINC0)) == 0)
        {
        }

        // Serial.println("ZeroCross true");

        // If received bit is 0 on Port C pin 0
        if ((PINC & (1 << PINC0)) != 0)
        {

            recvChar |= (1 << readBitIndex);
            // Serial.println("Received bit 1");
        }
        // if received bit != 0
        else
        {
            recvChar &= ~(1 << readBitIndex);
            // Serial.println("Received bit 0");
        }

        readBitIndex--;
        zerocross = false;
    }
    // Serial.println(receivedChar_);

    // Serial.println("--------------------");
    // Serial.println("--------------------");
    // Serial.println("--------------------");
    // Serial.println("--------------------");

    // Set startbit to 0 in receivedChar
    recvChar &= ~(1 << 4);
    return recvChar;
}

void X10 ::sendData(char c)
{
    while (bitIndex >= 0)
    {
        while (zerocross == false)
        {
            // do nothing
        }

        while (zerocross)
        {
            // Startbit tjek
            if ((c & (1 << 4)))
            {
                // do nothing
            }

            else
            { // set startbit to one
                c |= (1 << 4);
            }

            if (((c & (1 << bitIndex)) == 0))
            {
                // set pin 6 of port H low
                PORTH &= ~(1 << PH6);
                _delay_ms(1);
                bitIndex--;
                zerocross = false;
            }
            else
            {

                PORTH |= (1 << PH6);
                _delay_ms(1);
                PORTH &= ~(1 << PH6);

                // initPWM();
                bitIndex--;
                zerocross = false;
            }
        }
    }

    bitIndex = 4;
}
