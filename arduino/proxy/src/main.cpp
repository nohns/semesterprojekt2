
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
