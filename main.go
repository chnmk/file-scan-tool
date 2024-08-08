package main

import (
	"fmt"
	"os"

	"github.com/chnmk/file-scan-tool/cli"
	"github.com/chnmk/file-scan-tool/output"
	"github.com/chnmk/file-scan-tool/scanner"
)

func main() {
	fileFormat, err := cli.InputFileFormat(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Scanning for files:", fileFormat)
	result := scanner.ScanFiles(fileFormat)

	output.HandleResult(result)
}
