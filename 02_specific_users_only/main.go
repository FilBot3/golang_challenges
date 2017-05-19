package specific_person

import (
	"strings"
)

func Say_Hello_To_Specific_Person(name string) (result string) {
	switch {
	case strings.TrimRight(name, "\n") == "Phillip":
		result = "Hello Phillip"
        return result
	case name == "Phil\n":
		result = "Hello Phil"
        return result
	default:
		result = "You're not authorized to use this program."
        return result
	}
}
