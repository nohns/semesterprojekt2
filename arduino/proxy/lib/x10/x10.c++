#include <avr/io.h>
#include "x10.h"
#include <stdlib.h>


extern volatile bool zerocross;
volatile int flag=0;
volatile int bitIndex = 4;
volatile char receivedChar=0;
volatile static int receivedBit=0;
volatile bool dataHigh = false;




X10::X10()
{

}

char X10 :: readData() 
{
do{
while(zerocross=false)
{
    //do nothing just wait
}
    // set Timer 2 to generate interrupts at 120 kHz
TCCR2A = 0;  // set Timer 2 to normal mode
TCCR2B = (1 << CS20);  // set prescaler to 1
OCR2A = 83;  // set TOP value to 83 for 120 kHz frequency
TIMSK2 = (1 << OCIE2A);  // enable Timer 2 compare match interrupt

 while(flag==0){
    
//If received bit is 0
    if (PIND & (0<<PIND0)==(0<<PIND0))
    {
        receivedBit=0;
        
    }
    //if received bit != 0
    else{
        receivedBit=1;
        dataHigh=true;
    }
 }

 flag=0;

//set char bit value to received data
 if(dataHigh=true)
 {
    receivedChar|=(1<<bitIndex);
    dataHigh=false;
 }
 else{
    receivedChar|=~(1<<bitIndex);
 }


}while(bitIndex>=0);

 bitIndex=4;
 return (receivedChar& (1<<4) ? receivedChar : 0);
 
}


void X10 :: sendData(char c)
{
do
{
    while(zerocross=false)
{
    //do nothing just wait
}

char char_tx=c;
char_tx|=(1<<4);

 // Set Timer 1 to CTC mode and enable interrupt
TCCR1A = 0;
TCCR1B = (1 << WGM12) | (1 << CS12) | (1 << CS10);
OCR1A = 15;
TIMSK1 = (1 << OCIE1A);

while(flag==0)
{
 
    if((char_tx& (1<<bitIndex))==0 )
    {
        //set pin 6 of port H low 
        PORTH &= ~(1<<PH6);
    }
    else{
//sets pin 6 of port H to 120kHz
    //PWM at 120 kHz with DC=50% 
        //Timer 2 PWM fast mode enables and non-inverting mode is set and prescaler 8
        TCCR2A|10000011;
        TCCR2B|00001010;
        //TOP LEVEL set to 128 Timer 2
        OCR2A=127;
        //OCR value is set to 64 for Timer 2
        OCR2B=64;
    }
 }
    flag=0;

}while(bitIndex>=0)

bitIndex=4;

}


// Define timer ISR to be executed every 1 ms
ISR(TIMER1_COMPA_vect) {
    flag=1;
    bitIndex--;
    zerocross=false;
    // disable timer compare interrupt
    TIMSK1 &= ~(1 << OCIE1A);  
}

ISR(TIMER2_COMPA_vect)
{
    flag=1;
    zerocross=false;
    TCCR2B = 0; // disable Timer 2
}

