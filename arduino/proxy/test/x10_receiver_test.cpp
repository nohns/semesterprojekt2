#include "x10.h"
#include <stdio.h>
#include <stdlib.h>
#include <avr/interrupt.h>
#include "uart.h"


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

  char dataBuffer=0x00;

  X10 x10;
  Uart uart;

  while(true)
  {
    dataBuffer=x10.readData();
    uart.sendCommand(dataBuffer);
  }

  return 0;
}


//zerocross interrupt
 ISR(INT0_vect)
 {
  zerocross=true;
 }