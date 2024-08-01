package cli

import (
	"bytes"
	"testing"
)

func TestEmptyInput(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("\n"))

	fileFormat, err := InputFileFormat(&stdin)
	if err != nil {
		t.Error(err)
	}

	if fileFormat != ".txt" {
		t.Error("should be '.txt', got", fileFormat)
	}
}
