package acceptUserInput

import (
	"fmt"
)

func PrintName(name string) (result string) {
    result = fmt.Sprintf("Your name is: %s", name)
    return result
}