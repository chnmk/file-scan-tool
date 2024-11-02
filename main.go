package main

import (
	"fmt"

	"github.com/chnmk/file-scan-tool/cli"
	"github.com/chnmk/file-scan-tool/dialog"
	"github.com/chnmk/file-scan-tool/output"
	"github.com/chnmk/file-scan-tool/scanner"
	"github.com/ncruces/zenity"
)

func main() {
	var fileFormats []string
	var outputMode string

	cancelSwitchToCLI := zenity.Question("Switch to command line input?",
		zenity.Title("Question"),
		zenity.QuestionIcon)

	if cancelSwitchToCLI == nil {
		fileFormats, outputMode = cli.CliInit()
	} else {
		fileFormats, outputMode = dialog.DialogInit()
	}

	params := []string{"printParent"} // temp
	outputHandler := output.SelectHandler(outputMode, params)

	fmt.Println("Scanning for files:", fileFormats)
	result := scanner.ScanFiles(fileFormats)

	fmt.Println("Processing result...")
	outputHandler.HandleOutput(result)
}
