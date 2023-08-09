package makefile

type Rule struct {
	Targets      []string
	Dependencies []string
	Expressions  []string
}
