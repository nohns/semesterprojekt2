class Lock
{
private:
    // State of the lock
    bool isEngaged;

    // Pin number of the lock
    static const int pin = 1234;

public:
    // Constructor to create lock object
    Lock();

    // Get state of the lock
    bool getIsEngaged();

    // Change state of the lock
    void setIsEngaged(bool isEngaged);

    // Method to verify input pin vs stored pin
    bool verifyPin(int pin);
};