#include <stdio.h>
#include "libft.h"

void check_memset(char *str, char fill_char, size_t num) {
    ft_memset(str, fill_char, num);
    printf("str: %s", str);
}

int main(int argc, char *argv[]) {
    char str[13] = "Hello, World!";
    check_memset(str, *argv[argc-2], ft_atoi(argv[argc-1]));
    return 0;
}
