package structs

type Output struct {
	PrintParent bool
}

// Handles params that are used in every output struct
func (p *Output) HandleDefaultParams(params []string) {
	if params[0] == "printParent" {
		p.PrintParent = true
	} else {
		p.PrintParent = false
	}
}
