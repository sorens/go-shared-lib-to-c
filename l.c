#include <stdio.h>

#include "l.h"

int main() 
{
    printf("A simple c program that calls Go code...\n");
    GoString world = {"world", 5};
    printf(" => %s\n", Hello(world));
    GoString nurse = {"nurse", 5};
    printf(" => %s\n", Hello(nurse));
}
