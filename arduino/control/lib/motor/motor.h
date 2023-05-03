#pragma once

#include <Servo.h>

class MotorDriver
{
    // Class defined in the Servo library
    Servo myservo;
    int pos;

public:
    // Constructor to create motor object
    MotorDriver();

    void engageLock();
    void disengageLock();
};