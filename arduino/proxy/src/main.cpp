
#include "uart.h"
#include "controller.h"
#include "x10.h"
#include <stdio.h>
#include <stdlib.h>

volatile bool zerocross;


int main()
{
//PORTD as input
  DDRD|0x00;

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

  Controller controller(&x10);

  Uart uart(&controller);

  // Start uart eventHandler
  while (true)
  {
    uart.awaitRequest();
  }
  return 0;
}

//zerocross interrupt
 ISR(INT0_vect)
 {
  zerocross=true;
 }
