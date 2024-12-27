#include <stdio.h>
#include "libft.h"

#include "libft.h"

void	*ft_memmove(void *dst, const void *src, size_t len)
{
	char	*dst_cpy;
	char	*src_cpy;
	size_t	i;

	dst_cpy = (char *) dst;
	src_cpy = (char *) src;
	i = -1;
	if (len == 0 || dst == src)
		return (dst);
	if (dst < src)
	{
		while (++i < len)
		{
			dst_cpy[i] = src_cpy[i];
		}
	}
	else
	{
		while (len > 0)
		{
			dst_cpy[len - 1] = src_cpy[len - 1];
			len--;
		}
	}
	return (dst);
}
