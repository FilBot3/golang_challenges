// helloWorld package is a demonstration on how to use golang at is most basic
// level. This does not contain a main function meaning that it is supposed
// to be used as an imported package for a func main, or main.go in another
// package. 
//
// This package is also being used as a demonstration on how to properly 
// document golang packages, and code using different methods. 
// This method is inline with the code itself. As long as the comments are 
// directly touching, or against the line they're documenting, Golang will
// automatically associate the document, or comment block with that line. 
// A blank line in between will break that documentation. 
//
// However, a break with a comment and blank line, as above, will signafy
// that there is a new line, or a new paragraph in the documentation that 
// will be genreated. I'm using a column length of 80, but modern screens
// are much larger than 80 characters, and text editors are intelligent. 
// However, just to make it read a little bit better, and easier on the eyes
// I'm using 80 characters for the column limit. 
package helloWorld

import (
	"fmt"
)

// PrintHello displays the string "Hello World"
// This function takes no arguments, but returns a value that can be saved or
// printed directly from another fucntion. 
func PrintHello() (hello_world string) {
        // Since the variable hello_world is defined in the function header as
	// a string, we don't have to redefine it as a string, or redeclare it.
	// We can just assign value to it. 
	hello_world = "Hello World"
	// Print the value so we can see what it is. This isn't as useful in 
	// production level code, but since we're just doing exercises, it
	// can be useful in seeing what the code is doing. 
	fmt.Println(hello_world)
	// Return the variable hello_world and its value to whatever is calling
	// this function. 
	return hello_world
}
