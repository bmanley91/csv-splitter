package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"

	"manley.dev/csv-splitter/logger"
)

var l logger.Logger

type ParsedFile struct {
	header string
	lines  []string
}

func main() {
	lineCountPtr := flag.Int("lineCount", 500, "Maximum number of lines per output file.")
	inputFilePathPtr := flag.String("inputFilePath", "./input.csv", "Path to the CSV File that will be split.")
	outputDirPtr := flag.String("outputDir", "./output", "The directory in which the output should be written.")
	logEnabledPtr := flag.Bool("verbose", false, "Whether or not verbose logging is enabled")

	flag.Parse()

	fmt.Println("lineCount: ", *lineCountPtr)
	fmt.Println("inputFile: ", *inputFilePathPtr)

	l = logger.InitLogger(*logEnabledPtr)

	parsedFile, parseErr := readFile(*inputFilePathPtr)

	if parseErr != nil {
		l.Fatalf("Error reading/parsing file %v\n", parseErr)
	}

	l.Infof("Read file with header: '%v' and %d records.\n", parsedFile.header, len(parsedFile.lines))

	fileCount, writeErr := writeOutput(parsedFile, *outputDirPtr, *lineCountPtr)

	if writeErr != nil {
		l.Fatalf("Error writing output %v\n", writeErr)
	}

	l.Infof("Succesfully split input file into %d output files in %v.\n", fileCount, *outputDirPtr)
}

func readFile(path string) (parsedFile ParsedFile, err error) {
	file, openErr := os.Open(path)
	if openErr != nil {
		return parsedFile, openErr
	}
	defer file.Close()

	parsedFile, parseErr := parseFile(file)

	if parseErr != nil {
		return parsedFile, parseErr
	}

	return parsedFile, nil
}

func parseFile(file *os.File) (parsedFile ParsedFile, err error) {
	l.Infof("Parsing file %v\n", file.Name())
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return parsedFile, errors.New("no lines in file")
	}

	// Get the first line and set it as the header
	parsedFile.header = scanner.Text()

	// Add the rest of the lines to the lines array
	for scanner.Scan() {
		line := scanner.Text()
		parsedFile.lines = append(parsedFile.lines, line)
	}

	return parsedFile, nil
}

func writeOutput(parsedFile ParsedFile, outputDir string, maxLineCount int) (fileCount int, err error) {
	// Create output dir
	mkdirErr := os.Mkdir(outputDir, 0755)
	if mkdirErr != nil {
		return 0, mkdirErr
	}

	currentLineCount := 1
	fileCount = 1
	currentFile := createNewFileWithHeader(parsedFile.header, outputDir, fileCount)

	for _, line := range parsedFile.lines {
		if currentLineCount >= maxLineCount {
			// Close connection to current file
			currentFile.Close()

			// Increment number of files
			fileCount++

			// Create a new file
			currentFile = createNewFileWithHeader(parsedFile.header, outputDir, fileCount)

			// Reset current line count
			currentLineCount = 1
		}

		// Write current line to file
		writeLineToFile(currentFile, line)

		// Increment line count
		currentLineCount++
	}

	return fileCount, nil
}

func createNewFileWithHeader(header string, outputDir string, index int) os.File {
	l.Infof("Creating file #%d.\n", index)
	outputFilePath := fmt.Sprintf("%v/output_%d.csv", outputDir, index)
	file, err := os.Create(outputFilePath)
	if err != nil {
		l.Fatalf("Error creating file %v", err)
	}

	writeLineToFile(*file, header)

	return *file
}

func writeLineToFile(file os.File, line string) error {
	_, err := file.Write([]byte(fmt.Sprint(line, "\n")))
	return err
}
