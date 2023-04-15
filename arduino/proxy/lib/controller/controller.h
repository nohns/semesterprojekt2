#pragma once

#include "lock.h"

class Controller
{
private:
public:
    // constructor

    char *getLock(Lock lock);
    char *setLock(Lock lock);
};