package main

import (
	"fmt"
	"os"

	"github.com/chnmk/file-scan-tool/cli"
	"github.com/chnmk/file-scan-tool/dialog"
	"github.com/chnmk/file-scan-tool/output"
	"github.com/chnmk/file-scan-tool/scanner"
	"github.com/ncruces/zenity"
)

func main() {
	var fileFormats []string
	var outputMode string
	var err error

	cancelErr := zenity.Question("Switch to command line input?",
		zenity.Title("Question"),
		zenity.QuestionIcon)

	if cancelErr == nil {
		fileFormats, err = cli.InputFileFormat(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return
		}

		outputMode, err = cli.InputOutputMode(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return
		}

	} else {
		fileFormats, err = dialog.FileFormatDialog()
		if err != nil {
			fmt.Println(err)
			return
		}

		outputMode, err = dialog.OutputModeDialog()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	params := []string{"printParent"} // temp
	outputHandler := output.SelectHandler(outputMode, params)

	fmt.Println("Scanning for files:", fileFormats)
	result := scanner.ScanFiles(fileFormats)

	fmt.Println("Processing result...")
	outputHandler.HandleOutput(result)
}
