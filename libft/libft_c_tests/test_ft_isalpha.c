#include <stdio.h>
#include "libft.h"

int main(int argc, char **argv)
{
	(void)argc;
	printf("alpha: %d", ft_isalpha(argv[1][0]));
	return (0);
}
