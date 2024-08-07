package main

import (
	"fmt"
	"os"

	"github.com/chnmk/file-scan-tool/cli"
	"github.com/chnmk/file-scan-tool/output"
	"github.com/chnmk/file-scan-tool/scanner"
)

func main() {
	fileFormats, err := cli.InputFileFormat(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputMode := "print"             // temp
	params := []string{"printParent"} // temp
	outputHandler := output.SelectHandler(outputMode, params)

	fmt.Println("Scanning for files:", fileFormats)
	result := scanner.ScanFiles(fileFormats)

	fmt.Println("Processing result...")
	outputHandler.HandleOutput(result)

}
