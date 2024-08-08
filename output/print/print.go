package output

import "fmt"

type Print struct {
	Result [][2]string
}

func (p Print) HandleOutput() {
	for _, r := range p.Result {
		fmt.Printf("File: %s, Parent: \\%s\n", r[0], r[1])
	}
}
