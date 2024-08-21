package cli

import (
	"bytes"
	"testing"
)

func TestValidFileFormatInput(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("go, mp3, jpg\n"))

	fileFormats, err := InputFileFormat(&stdin)
	if err != nil {
		t.Error(err)
	}

	if fileFormats[0] != "go" {
		errorString := "should be 'go', got"
		t.Error(errorString, fileFormats)
	}

	if fileFormats[1] != "mp3" {
		errorString := "should be 'mp3', got"
		t.Error(errorString, fileFormats)
	}

	if fileFormats[2] != "jpg" {
		errorString := "should be 'jpg', got"
		t.Error(errorString, fileFormats)
	}
}

func TestEmptyFileFormatInput(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("\n"))

	fileFormats, err := InputFileFormat(&stdin)
	if err != nil {
		t.Error(err)
	}

	if fileFormats[0] != "mp3" {
		errorString := "should be 'go', got"
		t.Error(errorString, fileFormats)
	}

	if fileFormats[1] != "flac" {
		errorString := "should be 'flac', got"
		t.Error(errorString, fileFormats)
	}

	if fileFormats[2] != "ogg" {
		errorString := "should be 'ogg', got"
		t.Error(errorString, fileFormats)
	}
}

func TestValidOutputModeInput(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("print\n"))

	outputMode, err := InputOutputMode(&stdin)
	if err != nil {
		t.Error(err)
	}

	if outputMode != "print" {
		errorString := "should be 'print', got"
		t.Error(errorString, outputMode)
	}
}

func TestEmptyOutputModeInput(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("\n"))

	outputMode, err := InputOutputMode(&stdin)
	if err != nil {
		t.Error(err)
	}

	if outputMode != "print" {
		errorString := "should be 'print' (default), got"
		t.Error(errorString, outputMode)
	}
}
