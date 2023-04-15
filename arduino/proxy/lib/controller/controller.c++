
#include "uart.h"
#include "controller.h"

#include <avr/io.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

// Remember to free memory and perform null check
char *Controller::getLock(Lock lock)
{
    // Do some stuff with the lock object

    // Call x-10 to get the lock status

    // Return the lock status

    // Temporarily we are just going to return the request string

    // Allocate heap memory using new operator
    // char *idk = new char[10];

    int8_t len = snprintf(NULL, 0, "%s/%s", lock.id, lock.state + 1);
    char *str = (char *)malloc(sizeof(char) * (len + 1));
    if (str != NULL)
    {
        // malloced memory is not guaranteed to be zeroed out so we need to do it ourselves
        memset(str, 0, len);
        //  Copy the string to the dynamically allocated memory
        strcpy(str, lock.id);
        strcat(str, lock.state);
        // Return the dynamically allocated memory
        return str;
    }
    return NULL;
}

/*
 * This function returns the lock state for the given lock.
 *
 * @param: lock - a Lock object that contains the lock id and state.
 *
 * @return: a string containing the lock id and state, or NULL if there was an error.
 */

char *Controller::setLock(Lock lock)
{
    // Do some stuff with the lock object

    // Call x-10 to set the lock status

    // Return the lock status

    // Temporarily we are just going to return the request string

    int8_t len = snprintf(NULL, 0, "%s/%s", lock.id, lock.state + 1);
    char *str = (char *)malloc(sizeof(char) * (len + 1));
    if (str != NULL)
    {
        // malloced memory is not guaranteed to be zeroed out so we need to do it ourselves
        memset(str, 0, len);
        //  Copy the string to the dynamically allocated memory
        strcpy(str, lock.id);
        strcat(str, lock.state);
        // Return the dynamically allocated memory
        return str;
    }
    return NULL;
}