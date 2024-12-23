import os
import subprocess
import threading
import argparse
from collections import defaultdict

ExpectedRes = {
    "test_ft_memset.c": [
        "str: *****, World!",  
        "str: ###lo, World!",  
        "str: Hello, World!",  
        "str: $$$$$, World!",  
        "str: Wello, World!",  
    ],
    "test_ft_strlen.c": [
        "len: 13",  
        "len: 0",   
        "len: 5",   
        "len: 1",   
        "len: 500", 
    ],
    "test_ft_strncmp.c": [
        "cmp: 0",  
        "cmp: 1",  
        "cmp: -1", 
        "cmp: 0",  
        "cmp: 0",  
    ],
    "test_ft_atoi.c": [
        "num: 42",  
        "num: -42",  
        "num: 0",  
        "num: 2147483647",  
        "num: -2147483648",  
    ],
    "test_ft_isalpha.c": [
        "alpha: 1",  
        "alpha: 1",  
        "alpha: 0",  
        "alpha: 0",  
        "alpha: 0",  
    ],
    "test_ft_isdigit.c": [
        "digit: 1",  
        "digit: 1",  
        "digit: 1",  
        "digit: 0",  
        "digit: 0",  
    ],
    "test_ft_isalnum.c": [
        "alnum: 1",  
        "alnum: 1",  
        "alnum: 1",  
        "alnum: 0",  
        "alnum: 0",  
    ],
    "test_ft_calloc.c": [
        "alloc: [0, 0, 0, 0, 0]",  
        "alloc: []",  
        "alloc: [0, 0, ..., 0]",  
        "alloc: []",  
        "alloc: [0, 0, ..., 0]",  
    ],
    "test_ft_substr.c": [
        "substr: Hello",  
        "substr: World",  
        "substr: ",  
        "substr: Hello, World!",  
        "substr: Ab",  
    ],
    "test_ft_strjoin.c": [
        "join: HelloWorld",  
        "join: 42Abu",  
        "join: Test",  
        "join: Hello",  
        "join: ",  
    ],
    "test_ft_split.c": [
        "split: ['Hello', 'World']",  
        "split: ['One', 'Two']",  
        "split: ['42']",  
        "split: []",  
        "split: ['A', 'B', 'C']",  
    ],
    "test_ft_itoa.c": [
        "itoa: 42",  
        "itoa: -42",  
        "itoa: 0",  
        "itoa: 2147483647",  
        "itoa: -2147483648",  
    ],
    "test_ft_strtrim.c": [
        "trim: 'Hello'",  
        "trim: 'Abu'",  
        "trim: 'Hello'",  
        "trim: 'Hello'",  
        "trim: ''",  
    ],
    "test_ft_lstnew.c": [
        "list: [42]",  
        "list: [Hello]",  
        "list: [NULL]",  
        "list: [0]",  
        "list: [Test]",  
    ],
    "test_ft_lstadd_front.c": [
        "list: [2, 1]",  
        "list: [Hello]",  
        "list: [0, 42]",  
        "list: [Tail, Head]",  
        "list: [World, Hello]",  
    ],
    "test_ft_lstsize.c": [
        "size: 3",  
        "size: 0",  
        "size: 2",  
        "size: 1",  
        "size: 5",  
    ],
    "test_ft_lstadd_back.c": [
        "list: [1, 2]",  
        "list: [Hello]",  
        "list: [42, 0]",  
        "list: [Head, Tail]",  
        "list: [Hello, World]",  
    ],
    "test_ft_lstdelone.c": [
        "delete: Node1",  
        "delete: Node2",  
        "delete: NULL",  
        "delete: 42",  
        "delete: Test",  
    ],
}

InputArgs = {
    # Part 1 - Libc functions
    "test_ft_memset.c": [
        ["*", "5"],  # Replace 5 characters with '*'
        ["#", "3"],  # Replace 3 characters with '#'
        ["", "0"],   # No change
        ["$", "5"],  # Replace 5 characters with '$'
        ["W", "1"],  # Replace only 1 character with 'W'
    ],
    "test_ft_strlen.c": [
        ["Hello, World!"],  # Normal string
        [""],               # Empty string
        ["42Abu"],          # Short string
        ["A"],              # Single character
        ["a" * 500],        # Large string (using Python's string multiplication)
    ],
    "test_ft_strncmp.c": [
        ["abc", "abc", "3"],    # Same string
        ["abcd", "abcc", "4"],  # Different at position 3
        ["abc", "abcde", "3"],  # Compare only first 3 chars
        ["", "", "0"],         # Empty strings
        ["ABC", "abc", "3"],   # Case-sensitive comparison
    ],
    "test_ft_atoi.c": [
        ["42"],           # Positive number
        ["-42"],          # Negative number
        ["0"],            # Zero
        ["2147483647"],   # Maximum int
        ["-2147483648"],  # Minimum int
    ],
    "test_ft_isalpha.c": [
        ["A"],  # Uppercase letter
        ["z"],  # Lowercase letter
        ["1"],  # Non-alphabet character
        [" "],  # Space
        ["?"],  # Special character
    ],
    "test_ft_isdigit.c": [
        ["0"],  # Lower bound
        ["5"],  # Middle digit
        ["9"],  # Upper bound
        ["a"],  # Non-digit
        [" "],  # Space
    ],
    "test_ft_isalnum.c": [
        ["A"],  # Alphabet
        ["z"],  # Lowercase letter
        ["9"],  # Digit
        [" "],  # Space
        ["?"],  # Special character
    ],
    "test_ft_calloc.c": [
        ["5", "4"],         # Allocate 5 blocks of 4 bytes
        ["0", "10"],        # Allocate 0 blocks
        ["1", "100"],       # Allocate 1 block of 100 bytes
        ["10", "0"],        # Allocate 10 blocks of size 0
        ["1000", "1000"],   # Large allocation
    ],
    # Part 2 - Additional functions
    "test_ft_substr.c": [
        ["Hello, World!", "0", "5"],     # Extract "Hello"
        ["Hello, World!", "7", "5"],     # Extract "World"
        ["Hello, World!", "50", "10"],   # Out of bounds
        ["Hello, World!", "0", "50"],    # Length exceeds string size
        ["42Abu", "2", "2"],             # Extract "Ab"
    ],
    "test_ft_strjoin.c": [
        ["Hello", "World"],  # Join "Hello" and "World"
        ["42", "Abu"],       # Join "42" and "Abu"
        ["", "Test"],        # Join empty string and "Test"
        ["Hello", ""],       # Join "Hello" and empty string
        ["", ""],            # Join two empty strings
    ],
    "test_ft_split.c": [
        ["Hello World", " "],  # Split by space
        ["One;Two", ";"],      # Split by semicolon
        ["42", ","],           # No split
        ["", ","],             # Empty string
        ["A,B,C", ","],        # Split into 3 parts
    ],
    "test_ft_itoa.c": [
        ["42"],           # Positive number
        ["-42"],          # Negative number
        ["0"],            # Zero
        ["2147483647"],   # Maximum int
        ["-2147483648"],  # Minimum int
    ],
    "test_ft_strtrim.c": [
        ["  Hello  ", " "],   # Trim spaces from both ends
        ["42Abu42", "42"],    # Trim '4' and '2' from both ends
        ["Hello", "xyz"],     # No characters to trim
        ["***Hello***", "*"],  # Trim '*' from both ends
        ["", " "],             # Empty string
    ],
    # Bonus part (linked list functions)
    "test_ft_lstnew.c": [
        ["42"],        # Create a list node with value 42
        ["Hello"],     # Create a list node with value Hello
        ["NULL"],      # Create a list node with NULL value
        ["0"],         # Create a list node with 0
        ["Test"],      # Create a list node with value Test
    ],
    "test_ft_lstadd_front.c": [
        ["1", "2"],         # Add 2 to front of list containing 1
        ["NULL", "Hello"],   # Add Hello to front of empty list
        ["42", "0"],        # Add 0 to front of list containing 42
        ["Head", "Tail"],    # Add Tail to front of list containing Head
        ["Hello", "World"],  # Add World to front of list containing Hello
    ],
    "test_ft_lstsize.c": [
        ["1->2->3"],        # List with 3 elements
        ["NULL"],           # Empty list
        ["A->B"],           # List with 2 elements
        ["Z"],              # List with 1 element
        ["1->2->3->4->5"],  # List with 5 elements
    ],
    "test_ft_lstadd_back.c": [
        ["1", "2"],        # Add 2 to back of list containing 1
        ["NULL", "Hello"],  # Add Hello to back of empty list
        ["42", "0"],       # Add 0 to back of list containing 42
        ["Head", "Tail"],   # Add Tail to back of list containing Head
        ["Hello", "World"],  # Add World to back of list containing Hello
    ],
    "test_ft_lstdelone.c": [
        ["Node1"],  # Delete node with value Node1
        ["Node2"],  # Delete node with value Node2
        ["NULL"],   # Deleting a NULL node
        ["42"],     # Delete node with value 42
        ["Test"],   # Delete node with value Test
    ],
}


class TestResult:
    def __init__(self, filename, input_args, output, expected, passed):
        self.filename = filename
        self.input_args = input_args
        self.output = output
        self.expected = expected
        self.passed = passed


def compile_libft(path):
    """Compiles the libft.a using the Makefile in the specified path."""
    try:
        subprocess.run(["make", "-C", path], check=True)
    except subprocess.CalledProcessError as e:
        print(f"Error compiling libft: {e}")


def get_test_files(root):
    """Walks through the test directory and returns all C test files."""
    return [
        os.path.join(dirpath, file)
        for dirpath, _, filenames in os.walk(root)
        for file in filenames if file.endswith(".c")
    ]


def run_test(test_source, args):
    """Compiles and runs the C test binary with dynamic arguments."""
    binary_name = f"./output/{os.path.splitext(os.path.basename(test_source))[0]}"
    
    if not os.path.exists(binary_name):
        try:
            subprocess.run(["gcc", "-o", binary_name, test_source], check=True)
        except subprocess.CalledProcessError as e:
            return "", f"Failed to compile {binary_name}: {e}"
    
    try:
        output = subprocess.check_output([binary_name] + args, text=True)
        return output.strip(), None
    except subprocess.CalledProcessError as e:
        return e.output.strip(), f"Failed to execute {binary_name}: {e}"


def assert_contains(actual, expected):
    """Compares the actual and expected output and returns a boolean result."""
    return expected.strip() in actual.strip()


def cleanup():
    """Removes the compiled test binaries and recreates the output directory."""
    output_dir = "./output/"
    if os.path.exists(output_dir):
        for file in os.listdir(output_dir):
            file_path = os.path.join(output_dir, file)
            if os.path.isfile(file_path):
                os.remove(file_path)
    os.makedirs(output_dir, exist_ok=True)


def run_tests_for_file(file, inputs, grouped_results):
    """Runs tests for a single file and stores the results."""
    test_results = []
    for i, input_args in enumerate(inputs):
        output, error = run_test(file, input_args)
        expected = ExpectedRes.get(file, [])[i]  # Added get() to avoid KeyError
        print(f"Expected: {expected}, Got: {output}")
        passed = assert_contains(output, expected) if expected else False
        result = TestResult(file, input_args, output if not error else "Error", expected, passed)
        test_results.append(result)
    grouped_results[file] = test_results


def main():
    parser = argparse.ArgumentParser(description="Run libft C tests.")
    parser.add_argument("--path", type=str, default=os.getcwd(), help="Path to the libft project")
    parser.add_argument("-v", "--verbose", action="store_true", help="Enable verbose output")
    args = parser.parse_args()
    
    path = args.path
    verbose = args.verbose
    
    print("Compiling libft project...")
    compile_libft(path)
    
    test_files = get_test_files("./libft_c_tests/")
    grouped_results = defaultdict(list)
    threads = []
    
    for i,filePath in enumerate(test_files):
        file = os.path.basename(filePath)
        inputs = InputArgs.get(file, [])[i]  # Added get() to avoid KeyError
        print(f"Running tests for {file} with inputs: {inputs}")
        thread = threading.Thread(target=run_tests_for_file, args=(filePath, inputs, grouped_results))
        threads.append(thread)
        thread.start()
    
    for thread in threads:
        thread.join()
    
    for filename, results in grouped_results.items():
        print(f"\nTest Results for {filename}:")
        for result in results:
            status = "\033[32mPASS\033[0m" if result.passed else "\033[31mFAIL\033[0m"
            print(f"{status}: Test with args {result.input_args}: \033[33mExpected\033[0m: {result.expected} \033[34mGot\033[0m: {result.output}")
        print()  # Add a space between test groups
    
    cleanup()



if __name__ == "__main__":
    main()
