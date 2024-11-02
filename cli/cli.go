package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/chnmk/file-scan-tool/config"
)

func CliInit() ([]string, string) {
	fileFormats, err := InputFileFormat(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return []string{"error"}, "error"
	}

	outputMode, err := InputOutputMode(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return []string{"error"}, "error"
	}

	return fileFormats, outputMode
}

// Scans user input for comma-separated file format values
func InputFileFormat(stdin io.Reader) ([]string, error) {
	fmt.Println("Enter file formats separated by commas")
	fmt.Println("Or use '*' to scan all files")
	fmt.Printf("Default: %s\n", config.DefaultInput)
	reader := bufio.NewReader(stdin)

	// Read data
	fileFormat, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}

	result := trimFileFormat(fileFormat)

	return result, nil
}

// Converts user input to a slice
func trimFileFormat(fileFormat string) []string {
	// Trim spaces and check if input is empty
	fileFormat = strings.TrimSpace(fileFormat)
	if fileFormat == "" || fileFormat == "EOF" {
		fileFormat = config.DefaultInput
	}

	// Replace "," with " ", trim all punctuation
	fileFormat = strings.ReplaceAll(fileFormat, ",", " ")

	// Convert to []string separated with " "
	formatSlice := strings.Split(fileFormat, " ")

	// Remove empty strings from slice
	var formatSliceNoEmpty []string
	for _, f := range formatSlice {
		if f != "" {
			formatSliceNoEmpty = append(formatSliceNoEmpty, f)
		}
	}

	// Check if slice is empty
	if len(formatSliceNoEmpty) == 0 {
		fmt.Println("Input is empty, using default instead...")
		fileFormat = config.DefaultInput
	}

	return formatSliceNoEmpty
}

// Scans user input for output mode value
func InputOutputMode(stdin io.Reader) (string, error) {
	fmt.Printf("Enter output format (default: %s)\n", config.DefaultOutputMode)
	fmt.Println("print: prints output in the console")
	fmt.Printf("text: saves output in a text file (%s)\n", config.DefaultOutputFile)
	reader := bufio.NewReader(stdin)

	// Read data
	outputFormat, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	outputFormat = strings.TrimSpace(outputFormat)

	if outputFormat != "print" && outputFormat != "text" {
		fmt.Println("Input is invalid, using default instead...")
		outputFormat = config.DefaultOutputMode
	}

	return outputFormat, nil
}
