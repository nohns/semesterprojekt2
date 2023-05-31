#pragma once

class MotorDriver
{
private:
    static const int ICR_20MS_PERIOD = 4999;
    static const int OCR_LOCKED = 535;
    static const int OCR_UNLOCKED = 97;

public:
    // Constructor to create motor object
    MotorDriver();

    void engageLock();
    void disengageLock();
};