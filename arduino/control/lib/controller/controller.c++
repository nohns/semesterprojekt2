
#include "uart.h"
#include "controller.h"

#include <avr/io.h>
#include <Arduino.h>

// Request constants
const char OPEN_LOCK_CMD = 0b1011;
const char CLOSE_LOCK_CMD = 0b1010;
const char LOCK_STATE_CMD = 0b1000;

// Response constants
const char NOOP = 0b0000;     // No operation
const char ACK = 0b1111;      // Signals command being acknowledged
const char LOCKED = 0b1101;   // Signals lock engaged
const char UNLOCKED = 0b1110; // Signals lock disengages

Controller::Controller(MotorDriver *motor)
{
    this->motor = motor;

    // Create a lock object to contain the state
    this->lockState = Lock();

    // Make sure we synchronize the physical state of the lock, with the domain lock state
    if (this->lockState.getIsEngaged())
    {
        this->motor->engageLock();
    }
    else
    {
        this->motor->disengageLock();
    }
}

char Controller::routeCommand(char cmd)
{
    // Serial.println("cmd Routing");
    //  Route the command to the appropriate function
    switch (cmd)
    {
    case OPEN_LOCK_CMD:
        this->engageLock();
        return ACK;

    case CLOSE_LOCK_CMD:
        this->disengageLock();
        return ACK;

    case LOCK_STATE_CMD:
        bool state = this->getState();
        if (state)
        {
            return LOCKED;
        }
        else
        {
            return UNLOCKED;
        }
        break;
    }

    return NOOP;
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
    if (this->getState()))
    {
        // Call the motor to disengage the lock
        this->disengageLock();
    }
    else
    {
        // Call the motor to engage the lock
        this->engageLock();
    }

    return this->getState();
}

bool Controller::verifyPin(int pin)
{
    bool ok = lockState.verifyPin(pin);
    if (ok)
    {
        toggleLock();
    }

    // Do a comparison between the input pin and the stored pin
    return this->lockState.verifyPin(pin);
}
