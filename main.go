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
    fmt.Println(hw.Print_Hello())
    
    fmt.Print("Please enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
    fmt.Println(pn.Print_Name(name))
    fmt.Println(sp.Say_Hello_To_Specific_Person(name))
}