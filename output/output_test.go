package output

import (
	"testing"
)

func TestSelectHandlerPrint(t *testing.T) {
	outputHandler := SelectHandler("print", []string{"printParent"})

	if outputHandler == nil {
		t.Error("outputHandler should be not empty")
	}
}

func TestSelectHandlerInvalid(t *testing.T) {
	outputHandler := SelectHandler("invalid", []string{"printParent"})

	if outputHandler != nil {
		t.Error("outputHandler should be empty")
	}
}
