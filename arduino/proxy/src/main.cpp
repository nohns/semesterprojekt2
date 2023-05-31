#include "x10.h"
#include "controller.h"
#include "uart.h"
#include <stdio.h>
#include <stdlib.h>
#include <avr/interrupt.h>
#include <util/delay.h>
#include <Arduino.h>

volatile bool zerocross;

void initExInterrupt();

int main(void)
{

  //// PORTC  as input
  // DDRC = 0x00;

  // PORTH as output
  DDRH = 0xFF;

  initExInterrupt();

  // Serial.begin(9600);

  X10 x10;

  Controller Controller(&x10);

  Uart uart(&Controller);

  while (true)

  {
    uart.awaitRequest();
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
  /* PORTH |= (1 << PH6);
            _delay_ms(1);
            PORTH &= ~(1 << PH6); */
  zerocross = true;
}
