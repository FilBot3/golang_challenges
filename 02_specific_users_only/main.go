package specific_person

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Say_Hello_To_Specific_Person() {
	fmt.Print("Please enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	switch {
	case strings.TrimRight(name, "\n") == "Phillip":
		fmt.Println("Hello Phillip")
	case name == "Phil\n":
		fmt.Println("Hello Phil")
	default:
		fmt.Println("You're not authorized to use this program.")
	}
}
