package text

import (
	"fmt"
	"log"
	"os"

	"github.com/chnmk/file-scan-tool/config"
)

type Text struct {
	printParent bool
}

func (p *Text) HandleParams(params []string) {
	if params[0] == "printParent" {
		p.printParent = true
	} else {
		p.printParent = false
	}
}

func (p Text) HandleOutput(result [][2]string) {
	var lines []string

	for _, r := range result {
		if p.printParent {
			lines = append(lines, fmt.Sprintf("File: %s, Parent: %s\n", r[0], r[1]))
		} else {
			lines = append(lines, r[0])
		}
	}

	writeFile(lines)
}

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
