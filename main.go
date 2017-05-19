package main 

import (
    "bufio"
    "fmt"
    hw "github.com/predatorian3/golang_challenges/00_hello_world"
    pn "github.com/predatorian3/golang_challenges/01_accept_user_input"
    sp "github.com/predatorian3/golang_challenges/02_specific_users_only"
    "os"
)

func main() {
    fmt.Println(hw.PrintHello())
    
    fmt.Print("Please enter your name: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    name := scanner.Text()
    fmt.Println(name)
    
    for i,ch := range name {
		fmt.Printf("[%3d] %c %d\n", i,ch,ch)
	}

    fmt.Println(pn.PrintName(name))
    fmt.Println(sp.SayHelloToSpecificPerson(name))
}