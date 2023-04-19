#pragma once

#include "lock.h"
#include "x10.h"

class Controller
{
private:
public:
    // constructor
    Controller controller();

    char *getLock(Lock lock);
    char *setLock(Lock lock);
};