
#include "uart.h"
#include "controller.h"

#include <avr/io.h>

Controller::Controller(MotorDriver *motor)
{
    this->motor = motor;

    // Create a lock object to contain the state
    this->lockState = Lock();
}

void Controller::engageLock()
{

    // Call the motor to engage the lock
    this->motor->engageLock();

    // Update the state of the lock
    this->lockState.setIsEngaged(true);
}

void Controller::disengageLock()
{

    // Call the motor to disengage the lock
    this->motor->disengageLock();

    // Update the state of the lock
    this->lockState.setIsEngaged(false);
}

bool Controller::getState()
{

    return this->lockState.getIsEngaged();
}

bool Controller::toggleLock()
{
    // If the lock is engaged, disengage it
    if (this->lockState.getIsEngaged())
    {
        // Call the motor to disengage the lock
        this->motor->disengageLock();

        // Update the state of the lock
        this->lockState.setIsEngaged(false);
    }
    else
    {
        // Call the motor to engage the lock
        this->motor->engageLock();

        // Update the state of the lock
        this->lockState.setIsEngaged(true);
    }

    return this->lockState.getIsEngaged();
}