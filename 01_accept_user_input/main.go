package accept_user_input

import (
	"fmt"
)

func Print_Name(name string) (result string) {
    result = fmt.Sprintf("Your name is: %s", name)
    return result
}