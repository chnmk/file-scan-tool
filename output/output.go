package output

import (
	"fmt"

	p "github.com/chnmk/file-scan-tool/output/print"
	sql_api "github.com/chnmk/file-scan-tool/output/sql"
	"github.com/chnmk/file-scan-tool/output/text"
)

type OutputHandler interface {
	HandleOutput([][2]string)
}

// Selects output handling mode based on user input
func SelectHandler(outputMode string, params []string) OutputHandler {
	if outputMode == "print" {
		var handler p.Print
		handler.HandleParams(params)
		return handler
	}

	if outputMode == "text" {
		var handler text.Text
		handler.HandleParams(params)
		return handler
	}

	if outputMode == "sql" {
		var handler sql_api.SQLFile
		handler.HandleParams(params)
		return handler
	}

	/*
		if outputMode == "json" {
			// ...
			handler.HandleParams(params)
			return handler
		}

		if outputMode == "yaml" {
			// ...
			handler.HandleParams(params)
			return handler
		}

		if outputMode == "sql" {
			// ...
			handler.HandleParams(params)
			return handler
		}
	*/

	fmt.Println("Invalid output mode")
	return nil
}
