

#include "controller.h"

#include <avr/io.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

// Remember to free memory and perform null check

Controller::Controller(X10 *x10)
{
    this->x10 = x10;
}

char *Controller::forwardRequest(char *cmd)
{

    // char *res = x10.readData();

    // Call x-10 to send the command

    // Return the response from x-10

    // Temporarily we are just going to return the request string

    char locked[] = {0b1101};
    char unlocked[] = {0b1110};

    char *str = (char *)malloc(sizeof(char));
    if (str != NULL)
    {
        //  Copy the string to the dynamically allocated memory
        strcpy(str, unlocked);
        // Return the dynamically allocated memory
        return str;
    }
    return NULL;
}
