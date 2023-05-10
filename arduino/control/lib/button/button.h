#pragma once

#include "controller.h"

class Button
{

private:
    Controller *controller;

public:
    Button(Controller *controller);
    void isPressed();
};