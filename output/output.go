package output

import (
	"fmt"

	output "github.com/chnmk/file-scan-tool/output/print"
)

type OutputHandler interface {
	HandleOutput([][2]string)
}

func SelectHandler(outputMode string, params []string) OutputHandler {
	if outputMode == "print" {
		var handler output.Print
		handler.HandleParams(params)
		return handler
	}

	fmt.Println("Invalid output mode")
	return nil
}
