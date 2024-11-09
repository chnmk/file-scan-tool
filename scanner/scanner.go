package scanner

import (
	"io/fs"
	"log"
	"path/filepath"
)

// Walk the current directory searching for files of selected formats
func ScanFiles(fileFormats []string) [][2]string {
	var fileList [][2]string

	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Check if the format of the current file is in the list
		containsFormat := false
		for _, s := range fileFormats {
			if filepath.Ext(d.Name()) == "."+s {
				containsFormat = true
			}
		}

		// If the format is in the list, add it to the result array
		if containsFormat {
			var fileWithParent [2]string
			fileWithParent[1], fileWithParent[0] = filepath.Split(path)
			fileList = append(fileList, fileWithParent)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
		// return
	}

	return fileList
}
