package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Please enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is:", name)
}
