#pragma once

#include <Servo.h>

class MotorDriver
{
public:
    // Constructor to create motor object
    MotorDriver();

    void engageLock();
    void disengageLock();
};