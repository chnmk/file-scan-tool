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

	fmt.Println("Scanning for files:", fileFormat)
	// fileList := make(map[string]string)
	var fileList [][2]string
	result := scanner.ScanFiles(fileFormat, fileList)
	fmt.Println(result)
}
