#include "x10.h"
#include <stdio.h>
#include <stdlib.h>
#include <avr/interrupt.h>

volatile bool zerocross;
volatile bool flag=false;
volatile int bitIndex=4;

int main(void)
{

    //PORTD as input
  DDRC|0x00;

//PORTH as output
  DDRH|0xFF;


//enable all interrupts
  sei();

  //zero cross interrupt
  //external interrupt 0 activated
  EIMSK|0b00000001; 
  //interrupt when rising edge on int0 pin 0
  EICRA|0b00000011; 

  X10 x10;

  while(true)
  {
    x10.sendData(0b00001010);
  }

  return 0;
}


//zerocross interrupt
 ISR(INT0_vect)
 {
  zerocross=true;
 }
