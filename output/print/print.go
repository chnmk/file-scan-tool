package output

import "fmt"

func PrintResult(result [][2]string) {
	for _, r := range result {
		fmt.Printf("File: %s, Parent: \\%s\n", r[0], r[1])
	}
}
