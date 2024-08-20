package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/chnmk/file-scan-tool/config"
)

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
