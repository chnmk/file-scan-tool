package output

import (
	"fmt"

	output "github.com/chnmk/file-scan-tool/output/print"
)

type OutputHandler interface {
	HandleOutput()
}

func SelectHandler(outputMode string, result [][2]string) OutputHandler {
	if outputMode == "print" {
		handler := new(output.Print)
		handler.Result = result
		return handler
	}

	fmt.Println("Invalid output mode")
	return nil
}
