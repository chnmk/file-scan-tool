package cli

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/chnmk/file-scan-tool/config"
)

func TestEmptyInput(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("\n"))

	fileFormats, err := InputFileFormat(&stdin)
	if err != nil {
		t.Error(err)
	}

	if fileFormats[0] != config.DefaultInput {
		errorString := fmt.Sprintf("should be '%s' (default), got", config.DefaultInput)
		t.Error(errorString, fileFormats)
	}
}
