package output

import "fmt"

type Print struct {
	printParent bool
}

func (p *Print) HandleParams(params []string) {
	if params[0] == "printParent" {
		p.printParent = true
	} else {
		p.printParent = false
	}
}

func (p Print) HandleOutput(result [][2]string) {
	for _, r := range result {
		if p.printParent {
			fmt.Printf("File: %s, Parent: %s\n", r[0], r[1])
		} else {
			fmt.Printf("File: %s\n", r[0])
		}
	}
}
