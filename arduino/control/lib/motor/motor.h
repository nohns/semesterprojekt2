#pragma once

#include <../.pio/libdeps/megaatmega2560/Servo/src/Servo.h>

class MotorDriver
{
    Servo myservo;
    int pos;

public:
    // Constructor to create motor object
    MotorDriver();

    void engageLock();
    void disengageLock();
};