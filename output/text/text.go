package text

import (
	"fmt"
	"log"
	"os"

	"github.com/chnmk/file-scan-tool/config"
	"github.com/chnmk/file-scan-tool/output/structs"
)

type Text struct {
	structs.Output
}

/*
func (p *Text) HandleParams(params []string) {

}
*/

func (p Text) HandleOutput(result [][2]string) {
	var lines []string

	for _, r := range result {
		if p.PrintParent {
			lines = append(lines, fmt.Sprintf("File: %s, Parent: %s\n", r[0], r[1]))
		} else {
			lines = append(lines, r[0])
		}
	}

	writeFile(lines)
}

// Writes scanned file names to a text file
func writeFile(lines []string) {
	file, err := os.Create(config.DefaultOutputFile)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
