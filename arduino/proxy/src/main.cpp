/*
#include "uart.h"
#include "controller.h"
#include "x10.h"
#include <stdio.h>
#include <stdlib.h>
#include <avr/interrupt.h>

volatile bool zerocross;
volatile int bitIndex=5;

void initZerocrossInt();


int main()
{
//PORTC as input
  DDRC|0x00;

//PORTH as output
  DDRH|0xFF;


initZerocrossInt();
//enable all interrupts
  sei();

  //zero cross interrupt
  //external interrupt 0 activated
  EIMSK|0b00000001;
  //interrupt when rising edge on int0 pin 0
  EICRA|0b00000011;

  X10 x10;

  Controller controller(&x10);

  Uart uart(&controller);

  // Start uart eventHandler
  while (true)
  {

    x10.sendData(0b00001010);
    
  }
  return 0;
}

void initZerocrossInt()
{
  //zero cross interrupt
  //external interrupt 0 activated
  EIMSK|0b00000001; 
  //interrupt when rising edge on int0 pin 0
  EICRA|0b00000011; 
    //enable all interrupts
  sei();

}

//zerocross interrupt
 ISR(INT0_vect)
 {
  zerocross=true;
 }

 */

#include "x10.h"
#include <stdio.h>
#include <stdlib.h>
#include <avr/interrupt.h>
#include <util/delay.h>

volatile bool zerocross;
volatile int bitIndex = 5;

//void sendData(char);
void initExInterrupt();
//void exInterruptDisable();
//void PWMDisable();
//void initPWM();
//void initTimer1interrupt();

int main(void)
{
  /*
      //PORTC  as input
    DDRC=0x00;

    */

  // PORTH as output
  DDRH = 0xFF;

  initExInterrupt();

  X10 x10;

  while (true)
  {

    x10.sendData(0b00001010);
  }

  return 0;
}

void initExInterrupt()
{
  // zero cross INT0 enabled
  EIMSK |= 0b00000001;
  // interrupt when rising edge on int0 PD pin 0
  EICRA |= 0b00000011;
  // Globalt interrupt enable
  sei();
}

// zerocross interrut
ISR(INT0_vect)
{
  zerocross = true;
}

/* void sendData(char c)
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


        // Timer ellers når den aldrig længere
        //initPWM();
        bitIndex--;
        zerocross = false;
      }
    }
  }

  bitIndex = 5;
}

void initExInterrupt()
{
  // zero cross INT0 enabled
  EIMSK |= 0b00000001;
  // interrupt when rising edge on int0 PD pin 0
  EICRA |= 0b00000011;
  // Globalt interrupt enable
  sei();
}

// zerocross interrut
ISR(INT0_vect)
{
  zerocross = true;
}

void initPWM()
{
  initTimer1interrupt();
  TCCR2A = (1 << WGM21) | (1 << WGM20); // fast PWM
  TCCR2B = (1 << WGM22);                // fast PWM
  TCCR2A |= (1 << COM2B1);              // clear OC2B on Compare match
  OCR2A = 132;
  OCR2B = OCR2A / 2;
  TIMSK2 |= (1 << TOIE2); // timer 2 interrupt
  TCCR2B |= (1 << CS20);  // Start Timer/Counter2 with no prescaling
  sei();
}

void initTimer1interrupt()
{
  TCCR1A = 0;
  TCCR1B = (1 << WGM12) | (1 << CS11) | (1 << CS10); // Prescaler of 64
  OCR1A = 250;                                       // For 16 MHz CPU frequency and prescaler of 64, 250 gives a compare match after 1 ms
  TIMSK1 = (1 << OCIE1A);                            // Enable compare match interrupt for Timer/Counter1
  TCCR1B |= (1 << CS11) | (1 << CS10);               // Start Timer/Counter1 with prescaler of 64
}

ISR(TIMER1_COMPA_vect)
{
  TCCR2B = 0;           // Stop Timer/Counter2
  TIMSK1 = 0;           // Disable compare match interrupt for Timer/Counter1
  OCR2B = 0;            // Reset OCR2B value
  PORTH &= ~(1 << PH6); // Set pin 9 low
} */

/* #include "x10.h"
#include <stdio.h>
#include <stdlib.h>
#include <avr/interrupt.h>
#include <util/delay.h>

volatile bool zerocross;
volatile int bitIndex = 5;
volatile bool flag=false;
volatile char receivedChar;

char readData();
void initExInterrupt();
void exInterruptDisable();
void initPWM();


int main(void)
{
  /*
      //PORTC  as input
    DDRC=0x00;

    */

// PORTH as output
/*  DDRH = 0xFF;

 initExInterrupt();

 // X10 x10;

 while (true)
 {


 }

 return 0;
}

char readData()
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
           if ((PINC & (0 << PINC0)) == (0 << PINC0))
           {
               receivedChar |= ~(1 << bitIndex);
           }
           // if received bit != 0
           else
           {
               receivedChar |= (1 << bitIndex);
           }
     bitIndex--;
     zerocross=false;
   }

 }
  bitIndex=5;

  return receivedChar;
}


void initExInterrupt()
{
 // zero cross INT0 enabled
 EIMSK |= 0b00000001;
 // interrupt when rising edge on int0 PD pin 0
 EICRA |= 0b00000011;
 // Globalt interrupt enable
 sei();
}

// zerocross interrut
ISR(INT0_vect)
{
 zerocross = true;
}
*/
