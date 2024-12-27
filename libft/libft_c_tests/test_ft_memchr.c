#include <stdio.h>
#include "libft.h"

#include "libft.h"

void	*ft_memchr(const void *s, int c, size_t n)
{
	char	*s_copy;
	int		i;
	size_t	s_len;

	i = 0;
	s_copy = (char *)s;
	s_len = (size_t) ft_strlen(s_copy);
	if (n > s_len)
		n = s_len;
	while (i < (int)n)
	{
		if (s_copy[i] == c)
			return (&s_copy[i]);
		i++;
	}
	return (NULL);
}
