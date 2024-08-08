package scanner

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func ScanFiles(fileFormat string) [][2]string {
	var fileList [][2]string

	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(d.Name()) == ".go" {
			var parent string
			pathSlice := strings.Split(path, "\\")
			if len(pathSlice) > 1 {
				parent = pathSlice[len(pathSlice)-2] // -1 for slice index, -1 for prelast element
			}

			var fileWithParent [2]string
			fileWithParent[0] = d.Name()
			fileWithParent[1] = parent
			fileList = append(fileList, fileWithParent)
		}

		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return fileList
}
