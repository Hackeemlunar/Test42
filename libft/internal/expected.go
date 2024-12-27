package internal

// ExpectedRes is a map where the keys are test file names and the values are
var ExpectedRes = map[string][]string{
	// Part 1 - Libc functions
	"test_ft_memset.c": {
		"str: *****, World!", // Replace first 5 chars with '*'
		"str: ###lo, World!", // Replace first 3 chars with '#'
		"str: Hello, World!", // No change
		"str: $$$$$, World!", // Replace first 5 chars with '$'
		"str: Wello, World!", // Replace 'H' with 'W'
	},
	"test_ft_strlen.c": {
		"len: 13",  // Normal string
		"len: 0",   // Empty string
		"len: 5",   // Short string
		"len: 1",   // Single character
		"len: 500", // Large string
	},
	"test_ft_strncmp.c": {
		"strncmp: 0", // Same string
		"strncmp: 1", // First string is greater
		"strncmp: 0", // compare only first 3 chars
		"strncmp: 0", // Empty strings
		"strncmp: 1", // Case-sensitive comparison
	},
	"test_ft_atoi.c": {
		"num: 42",          // Positive number
		"num: -42",         // Negative number
		"num: 0",           // Zero
		"num: 2147483647",  // Maximum int value
		"num: -2147483648", // Minimum int value
	},
	"test_ft_isalpha.c": {
		"alpha: 1", // Uppercase letter (A)
		"alpha: 1", // Lowercase letter (z)
		"alpha: 0", // Non-alphabet character (1)
		"alpha: 0", // Space character
		"alpha: 0", // Special character (?)
	},
	"test_ft_isalnum.c": {
		"alnum: 1", // Alphabet (A)
		"alnum: 1", // Lowercase letter (z)
		"alnum: 1", // Digit (9)
		"alnum: 0", // Space
		"alnum: 0", // Special character (?)
	},
	"test_ft_isdigit.c": {
		"digit: 1", // '0' is a digit
		"digit: 1", // '5' is a digit
		"digit: 1", // '9' is a digit
		"digit: 0", // 'a' is not a digit
		"digit: 0", // Space is not a digit
	},
	"test_ft_calloc.c": {
		"alloc: [0, 0, 0, 0, 0]", // Allocate 5 blocks of 4 bytes (all zeros)
		"alloc: []",              // Allocate 0 blocks
		"alloc: [0, 0, ..., 0]",  // Allocate 1 block of 100 bytes
		"alloc: []",              // Allocate 10 blocks of size 0
		"alloc: [0, 0, ..., 0]",  // Large allocation
	},

	// Part 2 - Additional functions
	"test_ft_substr.c": {
		"substr: Hello",         // Extract "Hello"
		"substr: World",         // Extract "World"
		"substr: ",              // Out of bounds
		"substr: Hello, World!", // Full string when len exceeds size
		"substr: Ab",            // Extract "Ab"
	},
	"test_ft_strjoin.c": {
		"join: HelloWorld", // Join "Hello" and "World"
		"join: 42Abu",      // Join "42" and "Abu"
		"join: Test",       // Join empty string and "Test"
		"join: Hello",      // Join "Hello" and empty string
		"join: ",           // Join two empty strings
	},
	"test_ft_split.c": {
		"split: ['Hello', 'World']", // Split by space
		"split: ['One', 'Two']",     // Split by semicolon
		"split: ['42']",             // No split
		"split: []",                 // Empty string
		"split: ['A', 'B', 'C']",    // Multiple splits
	},
	"test_ft_itoa.c": {
		"itoa: 42",          // Positive number
		"itoa: -42",         // Negative number
		"itoa: 0",           // Zero
		"itoa: 2147483647",  // Maximum int
		"itoa: -2147483648", // Minimum int
	},
	"test_ft_strtrim.c": {
		"trim: 'Hello'", // Trim spaces from both ends
		"trim: 'Abu'",   // Trim '4' and '2' from both ends
		"trim: 'Hello'", // No characters to trim
		"trim: 'Hello'", // Trim '*' from both ends
		"trim: ''",      // Empty string
	},

	// Bonus part (linked list functions)
	"test_ft_lstnew.c": {
		"list: [42]",    // List with value 42
		"list: [Hello]", // List with value Hello
		"list: [NULL]",  // List with NULL value
		"list: [0]",     // List with value 0
		"list: [Test]",  // List with value Test
	},
	"test_ft_lstadd_front.c": {
		"list: [2, 1]",         // Add 2 to front of list containing 1
		"list: [Hello]",        // Add Hello to front of empty list
		"list: [0, 42]",        // Add 0 to front of list containing 42
		"list: [Tail, Head]",   // Add Tail to front of list containing Head
		"list: [World, Hello]", // Add World to front of list containing Hello
	},
	"test_ft_lstsize.c": {
		"size: 3", // List with 3 elements
		"size: 0", // Empty list
		"size: 2", // List with 2 elements
		"size: 1", // List with 1 element
		"size: 5", // List with 5 elements
	},
	"test_ft_lstadd_back.c": {
		"list: [1, 2]",         // Add 2 to back of list containing 1
		"list: [Hello]",        // Add Hello to back of empty list
		"list: [42, 0]",        // Add 0 to back of list containing 42
		"list: [Head, Tail]",   // Add Tail to back of list containing Head
		"list: [Hello, World]", // Add World to back of list containing Hello
	},
	"test_ft_lstdelone.c": {
		"delete: Node1", // Delete node with value Node1
		"delete: Node2", // Delete node with value Node2
		"delete: NULL",  // Deleting a NULL node
		"delete: 42",    // Delete node with value 42
		"delete: Test",  // Delete node with value Test
	},
}
