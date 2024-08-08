package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const defaultInput = "mp3, flac, ogg"

func InputFileFormat(stdin io.Reader) ([]string, error) {
	fmt.Println("Enter file formats separated by commas")
	fmt.Println("Or use '*' to scan all files")
	fmt.Printf("Default: %s\n", defaultInput)
	reader := bufio.NewReader(stdin)

	// Read data
	fileFormat, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}

	// Decompose this file into "read" and "trim"

	// Handle case when input is "EOF" (?)
	// ...

	// Trim spaces and check if input is empty
	fileFormat = strings.TrimSpace(fileFormat)
	if fileFormat == "" {
		fileFormat = defaultInput
	}

	// Replace "," with " ", trim all punctuation
	// ...

	// Convert to []string separated with " "
	// ...

	// Remove empty strings from slice
	// ...

	// Check if slice is empty
	// ...

	return []string{fileFormat}, nil
}
