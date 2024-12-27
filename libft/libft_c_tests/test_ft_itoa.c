#include <stdio.h>
#include "libft.h"

int main(int argc, char **argv)
{
	(void)argc;

	int n = ft_atoi(argv[1]);
	char *str = ft_itoa(n);
	if (str != NULL)
		printf("itoa: %s", str);
	else
		printf("itoa returned NULL");
	
	return (0);
}