package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func InputFileFormat(stdin io.Reader) (string, error) {
	fmt.Println("Enter file format (default: '.txt'):")
	reader := bufio.NewReader(stdin)

	fileFormat, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fileFormat = strings.TrimSpace(fileFormat)
	if fileFormat == "" {
		fileFormat = ".txt"
	}

	return fileFormat, nil
}
