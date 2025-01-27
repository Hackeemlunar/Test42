#include <stdio.h>
#include "libft.h"

#include "libft.h"

void ft_striteri(char *s, void (*f)(unsigned int, char *))
{
    unsigned int i = 0;
    while (s[i]) {
        f(i, &s[i]);
        i++;
    }
}
