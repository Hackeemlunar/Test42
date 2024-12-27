#include <stdio.h>
#include "libft.h"

#include "libft.h"

int main(int argc, char **argv)
{
	(void)argc;
	(void)argv;

	size_t count = 5;
	size_t size = 10;
	void *ptr = ft_calloc(count, size);
	if (ptr == NULL)
	{
		printf("ft_calloc returned NULL\n");
		return (1);
	}
	for (size_t i = 0; i < count * size; i++)
	{
		if (((unsigned char *)ptr)[i] != 0)
		{
			printf("ft_calloc did not zero-initialize the memory\n");
			return (1);
		}
	}
	printf("ft_calloc successfully zero-initialized the memory\n");
	return (0);
}