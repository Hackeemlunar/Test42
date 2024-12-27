#include <stdio.h>
#include "libft.h"

int main(int argc, char *argv[]) {
	char str[13] = "Hello, World!";
	ft_bzero(str, ft_atoi(argv[argc-1]));
	printf("str: %s", str);
	return 0;
}
