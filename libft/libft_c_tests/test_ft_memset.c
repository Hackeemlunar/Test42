#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void check_memset(char *str, char fill_char, size_t num) {
    memset(str, fill_char, num);
    printf("str: %s", str);
}

int main(int argc, char *argv[]) {
    char str[13] = "Hello, World!";
    check_memset(str, *argv[argc-2], atoi(argv[argc-1]));
    return 0;
}