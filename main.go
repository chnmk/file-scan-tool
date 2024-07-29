package main

import (
	"github.com/chnmk/file-scan-tool/cli"
	"github.com/chnmk/file-scan-tool/scanner"
)

func main() {
	fileFormat := cli.ReadFileFormat()
	scanner.ScanRecursive(fileFormat)
}
