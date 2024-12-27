#include <stdio.h>
#include "libft.h"

#include "libft.h"

size_t	ft_strlcat(char *dst, const char *src, size_t dstsize)
{
	unsigned int	i;
	unsigned int	src_size;
	unsigned int	dest_size;

	i = 0;
	src_size = ft_strlen(src);
	dest_size = ft_strlen(dst);
	if (dstsize <= dest_size || dstsize == 0)
		return (dstsize + src_size);
	while (src[i] != '\0' && i < dstsize - dest_size - 1)
	{
		dst[dest_size + i] = src[i];
		i++;
	}
	dst[dest_size + i] = '\0';
	return (dest_size + src_size);
}
