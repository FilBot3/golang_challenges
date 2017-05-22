package specificPerson

import (
	"fmt"
	"strings"
)

func SayHelloToSpecificPerson(name string) (result string) {
	switch {
	case strings.TrimRight(name, "\n") == "Phillip":
		result = "Hello Phillip"
	case name == "Phil":
		result = "Hello Phil"
	default:
		result = "You're not authorized to use this program."
	}

	fmt.Println(result)
	return result
}
