

#include "controller.h"

#include <avr/io.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

// Remember to free memory and perform null check

Controller::Controller(X10 *x10)
{
    this->x10 = x10;
}

char Controller::forwardRequest(char cmd)
{

    // char res = x10.readData();
    x10->sendData(cmd);

    char locked = 0b1101;
    char unlocked = 0b1100;

    return locked;
}
