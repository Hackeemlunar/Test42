#include <stdio.h>
#include "libft.h"

int main(int argc, char **argv)
{
	(void)argc;
	printf("digit: %d", ft_isdigit(argv[1][0]));
	return (0);
}

