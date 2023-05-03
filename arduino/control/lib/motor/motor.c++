#include "motor.h"

#include <util/delay.h>

// Constructor to create motor object
MotorDriver::MotorDriver()
{
    Servo myservo; // create servo object to control a servo

    myservo.attach(9); // attaches the servo on pin 9 to the servo object

    // variable to store the servo position
    int pos = 0;
}

void MotorDriver::engageLock()
{
    // Engage the lock
    for (pos = 0; pos <= 180; pos += 1)
    { // goes from 0 degrees to 180 degrees
        // in steps of 1 degree
        myservo.write(pos); // tell servo to go to position in variable 'pos'
                            // waits 15 ms for the servo to reach the position

        _delay_ms(15);
    }
}

void MotorDriver::disengageLock()
{
    // Disengage the lock
    for (pos = 180; pos >= 0; pos -= 1)
    {                       // goes from 180 degrees to 0 degrees
        myservo.write(pos); // tell servo to go to position in variable 'pos'
        _delay_ms(15);      // waits 15 ms for the servo to reach the position
    }
}
