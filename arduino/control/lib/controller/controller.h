#pragma once

// #include "x10.h"

#include "lock.h"
#include "motor.h"

class Controller
{
private:
    Lock lockState;

    MotorDriver *motor;

public:
    // The constuctor is responsible for initializing the Lock composition
    Controller(MotorDriver *motor);

    void engageLock();

    void disengageLock();

    bool getState();

    bool toggleLock();
};