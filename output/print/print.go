package output

import "fmt"

type Print struct {
	PrintParent bool
}

func (p Print) HandleOutput(result [][2]string) {
	for _, r := range result {
		if p.PrintParent {
			fmt.Printf("File: %s, Parent: \\%s\n", r[0], r[1])
		} else {
			fmt.Printf("File: %s\n", r[0])
		}
	}
}
