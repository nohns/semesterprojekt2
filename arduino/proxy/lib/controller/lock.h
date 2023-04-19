#pragma once

class Lock
{
private:
public:
    Lock(char *id, char *state)
    {
        this->id = id;
        this->state = state;
    }

    char *id;
    // bool state;
    char *state;
};
