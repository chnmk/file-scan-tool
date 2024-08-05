package cli

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEmptyInput(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("\n"))

	fileFormat, err := InputFileFormat(&stdin)
	if err != nil {
		t.Error(err)
	}

	if fileFormat != defaultInput {
		errorString := fmt.Sprintf("should be '%s' (default), got", defaultInput)
		t.Error(errorString, fileFormat)
	}
}
