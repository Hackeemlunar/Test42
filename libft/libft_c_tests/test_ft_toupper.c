#include <stdio.h>
#include "libft.h"

#include "libft.h"

int	ft_toupper(int c)
{
	if (c >= 'a' && c <= 'z')
		return ('A' + (c - 'a'));
	return (c);
}
