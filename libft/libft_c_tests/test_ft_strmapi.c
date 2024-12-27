#include <stdio.h>
#include "libft.h"

#include "libft.h"

char *ft_strmapi(char const *s, char (*f)(unsigned int, char))
{
    char *result;
    unsigned int i;

    if (!s || !f)
        return NULL;

    result = malloc((ft_strlen(s) + 1) * sizeof(char));
    if (!result)
        return NULL;
    i = 0;
    while (s[i] != '\0')
    {
        result[i] = f(i, s[i]);
        i++;
    }
    result[i] = '\0';
    return result;
}