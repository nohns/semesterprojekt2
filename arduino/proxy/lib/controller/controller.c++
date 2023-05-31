

#include "controller.h"
#include <avr/io.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <util/delay.h>


// Remember to free memory and perform null check

Controller::Controller(X10 *x10)
{
    this->x10 = x10;
}

char Controller::forwardRequest(char cmd)
{

    // char res = x10.readData();
    for (size_t i = 0; i < 10; i++)
    {
        x10->sendData(cmd);
        _delay_ms(10);
    }

    //return message from controller
    char ctrlreturn=0x00;

    ctrlreturn=x10->readData();

    return ctrlreturn;
}

