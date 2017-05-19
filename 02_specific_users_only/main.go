package specificPerson

import (
	"strings"
)

func SayHelloToSpecificPerson(name string) (result string) {
	switch {
	case strings.TrimRight(name, "\n") == "Phillip":
		result = "Hello Phillip"
		return result
	case name == "Phil":
		result = "Hello Phil"
		return result
	default:
		result = "You're not authorized to use this program."
		return result
	}
}
