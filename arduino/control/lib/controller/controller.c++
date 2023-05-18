
#include "uart.h"
#include "controller.h"

#include <avr/io.h>

// request constants
const char openLockCmd = 0b1011;
const char closeLockCmd = 0b1010;
const char lockStateCmd = 0b1000;

// Response constants
const char ack = 0b1111;
const char nack = 0b1100;
const char locked = 0b1101;
const char unlocked = 0b1110;

Controller::Controller(MotorDriver *motor)
{
    this->motor = motor;

    // Create a lock object to contain the state
    this->lockState = Lock();
}

char Controller::routeCommand(char cmd)
{
    // Route the command to the appropriate function
    switch (cmd)
    {
    case openLockCmd:
        this->engageLock();
        return ack;
        break;
    case closeLockCmd:
        this->disengageLock();
        return ack;
        break;
    case lockStateCmd:
        bool state = this->getState();
        if (state)
        {
            return locked;
        }
        else
        {
            return unlocked;
        }

        break;
    default:
        break;
    }
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