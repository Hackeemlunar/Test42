#include <stdio.h>
#include "libft.h"

int main(int argc, char **argv)
{
	(void)argc;
	printf("alnum: %d", ft_isalnum(argv[1][0]));
	return (0);
}
