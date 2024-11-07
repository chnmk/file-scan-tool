package output

import (
	p "github.com/chnmk/file-scan-tool/output/print"
	sql_api "github.com/chnmk/file-scan-tool/output/sql"
	"github.com/chnmk/file-scan-tool/output/text"
)

type OutputHandler interface {
	HandleOutput([][2]string)
	HandleDefaultParams([]string)
}

// Selects output handling mode based on user input
func SelectHandler(outputMode string, params []string) OutputHandler {
	var handler OutputHandler
	switch outputMode {
	case "print":
		handler = new(p.Print)
	case "text":
		handler = new(text.Text)
	case "sql":
		handler = new(sql_api.SQLFile)
	default:
		handler = new(p.Print)
	}

	handler.HandleDefaultParams(params)
	return handler
}
