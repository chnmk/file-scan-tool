package dialog

import (
	"github.com/ncruces/zenity"
)

// Placeholder
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
