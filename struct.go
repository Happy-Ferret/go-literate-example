// Initialize package
package main

// Imports
import "fmt"

// Declare struct to hold text field and method.
type Hello struct {
	Text string
}

// Declare struct method
func (a Hello) hello() {
	fmt.Println(a.Text)
}

func main() {
	// Initialize struct
	a := Hello {
		Text: "Hello World",
	}
	// Call struct method
	a.hello()
}
