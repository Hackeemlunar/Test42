package internal

import "strings"

// InputData is a map where the keys are test file names and the values are
var InputData = map[string][][]string{
	// Part 1 - Libc functions
	"test_ft_memset.c": {
		{"*", "5"}, // Replace 5 characters with '*'
		{"#", "3"}, // Replace 3 characters with '#'
		{"", "0"},  // No change
		{"$", "5"}, // Replace 5 characters with '$'
		{"W", "1"}, // Replace only 1 character with 'W'
	},
	"test_ft_strlen.c": {
		{"Hello, World!"},          // Normal string
		{""},                       // Empty string
		{"42Abu"},                  // Short string
		{"A"},                      // Single character
		{strings.Repeat("a", 500)}, // Large string
	},
	"test_ft_strncmp.c": {
		{"abc", "abc", "3"},   // Same string
		{"abcd", "abcc", "4"}, // Different at position 3
		{"abc", "abcde", "3"}, // Compare only first 3 chars
		{"", "", "0"},         // Empty strings
		{"ABC", "abc", "3"},   // Case-sensitive comparison
	},
	"test_ft_atoi.c": {
		{"42"},          // Positive number
		{"-42"},         // Negative number
		{"0"},           // Zero
		{"2147483647"},  // Maximum int
		{"-2147483648"}, // Minimum int
	},
	"test_ft_isalpha.c": {
		{"A"}, // Uppercase letter
		{"z"}, // Lowercase letter
		{"1"}, // Non-alphabet character
		{" "}, // Space
		{"?"}, // Special character
	},
	"test_ft_isdigit.c": {
		{"0"}, // Lower bound
		{"5"}, // Middle digit
		{"9"}, // Upper bound
		{"a"}, // Non-digit
		{" "}, // Space
	},
	"test_ft_isalnum.c": {
		{"A"}, // Alphabet
		{"z"}, // Lowercase letter
		{"9"}, // Digit
		{" "}, // Space
		{"?"}, // Special character
	},
	"test_ft_calloc.c": {
		{"5", "4"},       // Allocate 5 blocks of 4 bytes
		{"0", "10"},      // Allocate 0 blocks
		{"1", "100"},     // Allocate 1 block of 100 bytes
		{"10", "0"},      // Allocate 10 blocks of size 0
		{"1000", "1000"}, // Large allocation
	},

	// Part 2 - Additional functions
	"test_ft_substr.c": {
		{"Hello, World!", "0", "5"},   // Extract "Hello"
		{"Hello, World!", "7", "5"},   // Extract "World"
		{"Hello, World!", "50", "10"}, // Out of bounds
		{"Hello, World!", "0", "50"},  // Length exceeds string size
		{"42Abu", "2", "2"},           // Extract "Ab"
	},
	"test_ft_strjoin.c": {
		{"Hello", "World"}, // Join "Hello" and "World"
		{"42", "Abu"},      // Join "42" and "Abu"
		{"", "Test"},       // Join empty string and "Test"
		{"Hello", ""},      // Join "Hello" and empty string
		{"", ""},           // Join two empty strings
	},
	"test_ft_split.c": {
		{"Hello World", " "}, // Split by space
		{"One;Two", ";"},     // Split by semicolon
		{"42", ","},          // No split
		{"", ","},            // Empty string
		{"A,B,C", ","},       // Split into 3 parts
	},
	"test_ft_itoa.c": {
		{"42"},          // Positive number
		{"-42"},         // Negative number
		{"0"},           // Zero
		{"2147483647"},  // Maximum int
		{"-2147483648"}, // Minimum int
	},
	"test_ft_strtrim.c": {
		{"  Hello  ", " "},   // Trim spaces from both ends
		{"42Abu42", "42"},    // Trim '4' and '2' from both ends
		{"Hello", "xyz"},     // No characters to trim
		{"***Hello***", "*"}, // Trim '*' from both ends
		{"", " "},            // Empty string
	},

	// Bonus part (linked list functions)
	"test_ft_lstnew.c": {
		{"42"},    // Create a list node with value 42
		{"Hello"}, // Create a list node with value Hello
		{"NULL"},  // Create a list node with NULL value
		{"0"},     // Create a list node with 0
		{"Test"},  // Create a list node with value Test
	},
	"test_ft_lstadd_front.c": {
		{"1", "2"},         // Add 2 to front of list containing 1
		{"NULL", "Hello"},  // Add Hello to front of empty list
		{"42", "0"},        // Add 0 to front of list containing 42
		{"Head", "Tail"},   // Add Tail to front of list containing Head
		{"Hello", "World"}, // Add World to front of list containing Hello
	},
	"test_ft_lstsize.c": {
		{"1->2->3"},       // List with 3 elements
		{"NULL"},          // Empty list
		{"A->B"},          // List with 2 elements
		{"Z"},             // List with 1 element
		{"1->2->3->4->5"}, // List with 5 elements
	},
	"test_ft_lstadd_back.c": {
		{"1", "2"},         // Add 2 to back of list containing 1
		{"NULL", "Hello"},  // Add Hello to back of empty list
		{"42", "0"},        // Add 0 to back of list containing 42
		{"Head", "Tail"},   // Add Tail to back of list containing Head
		{"Hello", "World"}, // Add World to back of list containing Hello
	},
	"test_ft_lstdelone.c": {
		{"Node1"}, // Delete node with value Node1
		{"Node2"}, // Delete node with value Node2
		{"NULL"},  // Deleting a NULL node
		{"42"},    // Delete node with value 42
		{"Test"},  // Delete node with value Test
	},
}
