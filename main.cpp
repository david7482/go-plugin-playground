#include "calc.h"
#include <stdio.h>
#include <string.h>

int main() {
    printf("This is a C Application.\n");

    GoString name;
    name.p = "World";
    name.n = strlen(name.p);
    SayHello(name);

    printf("Add(3, 5): %d\n", Add(3, 5));
    return 0;
}
