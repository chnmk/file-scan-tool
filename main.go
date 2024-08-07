package main

import (
	"fmt"
	"os"

	"github.com/chnmk/file-scan-tool/cli"
	"github.com/chnmk/file-scan-tool/scanner"
)

func main() {
	fileFormat, err := cli.InputFileFormat(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	var fileList [][2]string

	fmt.Println("Scanning for files:", fileFormat)
	result := scanner.ScanFiles(fileFormat, fileList)

	for _, r := range result {
		fmt.Printf("File: %s, Parent: \\%s\n", r[0], r[1])
	}
}
