#pragma once

#include "x10.h"

class Controller
{
private:
    X10 *x10;

public:
    // constructor
    Controller(X10 *x10);

    char forwardRequest(char);
};