
#include "uart.h"
#include "controller.h"
// #include "x10.h"

#include <avr/io.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

Controller::Controller(MotorDriver *motor)
{
    this->motor = motor;

    // Create a lock object to contain the state
    this->lockState = Lock();
}

void Controller::engageLock()
{
}

void Controller::disengageLock()
{
}

bool Controller::getState()
{

    return this->lockState.getIsEngaged();
}

bool Controller::toggleLock()
{
}