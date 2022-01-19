# CSV Splitter
This is a command line tool used to split large CSVs into smaller files based on a max line count.

## Building
Running the following will build an executable called `csv-splitter` in this project's base directory.
```sh
$ go build
```

## Testing
Running the following will execute all tests.
```sh
$ go test
```

## Usage
This script supports the following flags:
* inputFilePath `string`- Path to the CSV File that will be split. (default "./input.csv")
* lineCount `int`- Maximum number of lines per output file. (default 500)
* outputDir `string`- The directory in which the output should be written. (default "./output")

Flags can be specified like so:
### Unix
```sh
$ ./csv-splitter \
    -inputFilePath /Users/me/path/to/file \
    -lineCount 100 \
    -outputDir ./output-goes-here
```
### Windows
```cmd
> csv-splitter.exe ^ 
      -inputFilePath \path\to\output ^
      -lineCount 100 ^
      -outputDir ./output-goes-here
```


Running the executable with no flags will use the defaults listed above.

Help can be displayed with the `-h` flag.
```sh
$ ./csv-splitter -h
Usage of ./csv-splitter:
  -inputFilePath string
        Path to the CSV File that will be split. (default "./input.csv")
  -lineCount int
        Maximum number of lines per output file (default 500)
  -outputDir string
```