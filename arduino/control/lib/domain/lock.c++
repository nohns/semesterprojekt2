#include "lock.h"

// Constructor to create lock object
Lock::Lock()
{
    isEngaged = false;
}

// Get state of the lock
bool Lock::getIsEngaged()
{
    return isEngaged;
}

// Change state of the lock
void Lock::setIsEngaged(bool isEngaged)
{
    this->isEngaged = isEngaged;
}

// Method to verify input pin vs stored pin
bool Lock::verifyPin(int pin)
{
    // Do a comparison between the input pin and the stored pin
    return this->pin == pin;
}