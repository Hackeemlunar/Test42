#include <stdio.h>
#include "libft.h"

int main(int argc, char **argv)
{
	(void)argc;
	printf("strncmp: %d", ft_strncmp(argv[1], argv[2], ft_atoi(argv[3])));
	return (0);
}
