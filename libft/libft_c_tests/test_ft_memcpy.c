#include <stdio.h>
#include "libft.h"

#include "libft.h"

void	*ft_memcpy(void *dst, const void *src, size_t n)
{
	int		i;
	char	*dst_cpy;
	char	*src_cpy;

	i = 0;
	dst_cpy = (char *) dst;
	src_cpy = (char *) src;
	while (i < (int) n)
	{
		dst_cpy[i] = src_cpy[i];
		i++;
	}
	return (dst);
}
