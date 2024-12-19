#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]) {
    printf("len: %zu", strlen(argv[argc-1]));
    return 0;
}