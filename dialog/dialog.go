package dialog

import (
	"github.com/ncruces/zenity"
)

// Placeholder file format select dialog
func FileFormatDialog() ([]string, error) {
	items, err := zenity.ListMultipleItems(
		"Select file formats:",
		"mp3", "flac", "ogg", "go",
	)

	if err != nil {
		return []string{}, err
	}

	return items, nil
}

// Output mode select dialog
func OutputModeDialog() (string, error) {
	mode, err := zenity.List(
		"Select output mode:",
		[]string{"print", "text"},
		zenity.Title("Select items from the list"),
		zenity.DisallowEmpty(),
	)

	if err != nil {
		return "print", err
	}

	return mode, nil
}
