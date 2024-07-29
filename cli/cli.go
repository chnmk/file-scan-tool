package cli

import (
	"fmt"
)

func ReadFileFormat() string {
	fileFormat := ".txt"

	fmt.Println("Enter file format (default: '.txt'):")
	fmt.Scan(&fileFormat)

	return fileFormat
}
