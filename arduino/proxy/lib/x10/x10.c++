#include <avr/io.h>
#include "x10.h"
#include <stdlib.h>
#include <avr/interrupt.h>
#include <util/delay.h>

extern volatile bool zerocross;
extern volatile int bitIndex;

X10::X10()
{

}

char X10 ::readData()
{
   while(bitIndex>=0)
 {

   while (zerocross == false)
       {
           // do nothing
       }

   while(zerocross)
   {

      // If received bit is 0 on Port C pin 0
           if ((PINC & (1 << PINC0)) != 0)
           {

            receivedChar_ |= (1 << bitIndex);
            
           }
           // if received bit != 0
           else
           {
               receivedChar_ &= ~(1 << bitIndex);
           }
           
     bitIndex--;
     zerocross=false;
   }

 }
  bitIndex=5;

  return receivedChar_;
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


        
        //initPWM();
        bitIndex--;
        zerocross = false;
      }
    }
  }

  bitIndex = 5;
}
