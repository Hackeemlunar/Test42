package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Hackeemlunar/42projects/internal"
)

const cTestRoot = "./libft_c_tests/"

type testResult struct {
	filename  string
	inputArgs []string
	output    string
	expected  string
	passed    bool
}

func main() {
	verbose := true
	// Parse the path and verbosity flags
	defaultPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	path := flag.String("path", defaultPath, "Path to the libft project")
	flag.BoolVar(&verbose, "v", false, "Enable verbose output")
	flag.Parse()

	// Compile the libft project
	log.Println("Compiling libft project...")
	if err := compileLibft(*path); err != nil {
		log.Printf("Error compiling libft: %v", err)
	}

	// Collect test files
	testFiles, err := getTestFiles(cTestRoot)
	if err != nil {
		log.Fatalf("Error finding test files: %v", err)
	}

	// Results grouped by test file
	groupedResults := make(map[string][]testResult)

	// WaitGroup to ensure all tests finish running
	var wg sync.WaitGroup

	for _, file := range testFiles {
		inputs := internal.InputData[filepath.Base(file)]
		wg.Add(1)
		go func(filename string, inputs [][]string) {
			defer wg.Done()
			var testResults []testResult
			for i, inputArgs := range inputs {
				result := testFtFunc(filename, inputArgs, i)
				testResults = append(testResults, result)
			}
			groupedResults[filename] = testResults
		}(file, inputs)
	}

	wg.Wait()

	// Print grouped results
	for filename, results := range groupedResults {
		fmt.Printf("\nTest Results for %s:\n", filename)
		for _, result := range results {
			if result.passed {
				fmt.Printf("\033[32mPASS\033[0m: ")
			} else {
				fmt.Printf("\033[31mFAIL\033[0m: ")
			}
			fmt.Printf("Test with args %v: \033[33mExpected\033[0m: %s \033[34mGot\033[0m: %s\033[0m\n",
				result.inputArgs, result.expected, result.output)
		}
		fmt.Println() // Add a space between test groups
	}

	cleanup() // Remove compiled test binaries
}

// getTestFiles walks through the test directory and returns all C test files
func getTestFiles(root string) ([]string, error) {
	var testFiles []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".c") {
			testFiles = append(testFiles, path)
		}
		return nil
	})
	return testFiles, err
}

// testFtFunc runs a single test case for a given file and returns a testResult
func testFtFunc(filename string, args []string, indx int) testResult {
	output, err := runTest(filename, args)
	if err != nil {
		log.Printf("Error running test %s: %v", filename, err)
		return testResult{filename, args, "Error running test", "", false}
	}
	expected := internal.ExpectedRes[filepath.Base(filename)][indx]
	passed := assertContains(output, expected)
	return testResult{filename, args, output, expected, passed}
}

// compileLibft compiles the libft.a using the Makefile in the specified path
func compileLibft(path string) error {
	cmd := exec.Command("make", "-C", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// runTest compiles and runs the C test binary with dynamic arguments
func runTest(testSource string, args []string) (string, error) {
	binaryName := "./output/" + strings.TrimSuffix(filepath.Base(testSource), ".c")

	// Check if the binary exists, compile it if not
	if _, err := os.Stat(binaryName); os.IsNotExist(err) {
		cmd := exec.Command("gcc", "-o", binaryName, testSource, "-L.", "-lft", "-I.")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return string(output), fmt.Errorf("failed to compile %s: %v", binaryName, err)
		}
	}

	// Run the compiled binary with the provided arguments
	runCmd := exec.Command(binaryName, args...)
	output, err := runCmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("failed to execute %s: %v", binaryName, err)
	}
	return string(output), nil
}

// cleanup removes the compiled test binaries
func cleanup() error {
	err := os.RemoveAll("./output/")
	if err != nil {
		return err // Handle the error appropriately
	}

	err = os.MkdirAll("./output/", 0755) // Recreate the directory
	if err != nil {
		return err // Handle the error appropriately
	}
	return nil
}

// assertContains compares the actual and expected output and returns a boolean result
func assertContains(actual, expected string) bool {
	actual = strings.TrimSpace(actual)
	expected = strings.TrimSpace(expected)
	return actual == expected
}
