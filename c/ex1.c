#include <stdio.h>

int main (int argc, char *argv[]) {
    int age = 22;
    float height = 1.73;
    printf("I am %d years old.", age);
    printf("I am %'.2f meters tall.", height);
    return 0;
}
