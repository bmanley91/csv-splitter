package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

const HEADER_TEXT = "first field, second field"
const BASE_LINE_TEXT = "line, number "

func TestReadFile(t *testing.T) {
	// Given an input file
	expectedParsedFile := createParsedFileWithLines(2)
	testFile := createTestFile(expectedParsedFile)
	defer os.Remove(testFile.Name())

	// When the file is read
	result, _ := readFile(testFile.Name())

	// Then it is properly parsed
	assertStringEquals(result.header, expectedParsedFile.header, t)
	assertStringEquals(result.lines[0], expectedParsedFile.lines[0], t)
	assertStringEquals(result.lines[1], expectedParsedFile.lines[1], t)
	assertIntEquals(len(result.lines), len(expectedParsedFile.lines), t)
}

func TestWriteOutput(t *testing.T) {
	// Given a desired max number of lines
	maxLines := 5

	// And a parsed input file with more than the max number of lines
	inputFileLines := maxLines
	parsedFile := createParsedFileWithLines(inputFileLines)

	// When output is generated
	outputDir := fmt.Sprint("test-output-", time.Now())
	outputFileCount, _ := writeOutput(parsedFile, outputDir, maxLines)

	// Then correct number of files are output
	assertIntEquals(outputFileCount, 2, t)

	// And the first file contains the max number of lines
	assertIntEquals(
		countFileLines(fmt.Sprint(outputDir, "/output_1.csv")),
		maxLines,
		t,
	)

	// And the second file contains the remaining lines
	assertIntEquals(
		countFileLines(fmt.Sprint(outputDir, "/output_2.csv")),
		2, // Header + One leftover row
		t,
	)
}

func createParsedFileWithLines(lineCount int) (output ParsedFile) {
	var lines []string
	index := 1
	for index <= lineCount {
		lines = append(lines, fmt.Sprint(BASE_LINE_TEXT, index))
		index++
	}

	output.header = HEADER_TEXT
	output.lines = lines

	return output
}

func createTestFile(parsedFile ParsedFile) *os.File {
	testFile, _ := os.CreateTemp("", "test-file.csv")

	testFile.Write([]byte(fmt.Sprint(parsedFile.header, "\n")))
	for _, line := range parsedFile.lines {
		testFile.Write([]byte(fmt.Sprint(line, "\n")))
	}

	testFile.Close()

	return testFile
}

func countFileLines(path string) (output int) {
	file, _ := os.Open(path)
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		output++
	}

	return output
}

func assertStringEquals(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("Incorrect answer. Expected %v, got %v", expected, actual)
	}
}

func assertIntEquals(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Errorf("Incorrect answer. Expected %d, got %d", expected, actual)
	}
}
