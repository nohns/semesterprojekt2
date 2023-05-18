#pragma once

// #include "x10.h"

#include "lock.h"
#include "motor.h"

class Controller
{
private:
    Lock lockState;

    MotorDriver *motor;

    void engageLock();

    void disengageLock();

    bool getState();

public:
    // The constuctor is responsible for initializing the Lock composition
    Controller(MotorDriver *motor);

    char routeCommand(char cmd);

    bool toggleLock();

    bool verifyPin(int pin);
};