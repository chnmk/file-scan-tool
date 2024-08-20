package scanner

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func ScanFiles(fileFormats []string) [][2]string {
	var fileList [][2]string

	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		containsFormat := false
		for _, s := range fileFormats {
			if filepath.Ext(d.Name()) == "."+s {
				containsFormat = true
			}
		}

		if containsFormat {
			var parent string
			pathSlice := strings.Split(path, "\\")
			if len(pathSlice) > 1 {
				parent = pathSlice[len(pathSlice)-2] // -1 to adjust for slice index, -1 for second to last element
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
