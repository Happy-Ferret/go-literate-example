// [Go](https://golang.org) is a modern programming language with a C-like syntax focused on concurrency, modularity and ease-of-use.

// [Hello World](https://en.wikipedia.org/wiki/%22Hello,_World!%22_program#Go) is a basic programming exercise done to showcase a programming language's syntax.

// ## Package definition
// Every executable uses the package name **`main`** as an entry point. 
package main

// ## Imports
// Here we import the **`fmt`** (short for "format") module. <br/>
// **`fmt`** contains simple routines to format and output text.
import "fmt"

// The **`Hello`** struct is a slight complication over simply running **`fmt.Println("Hello World")`**. <br/>
// This is done to showcase structs. One of the defining features of the Go programming language.

// While Go itself is not Object-Oriented, structs can act similar to classes by binding data and methods together.

// **`Text`** provides a simple string type.
type Hello struct {
	Text string
}

// **`hello()`** implements the **`Hello`** type struct. <br/>
// It returns **`Text`**.
func (a Hello) hello() {
	fmt.Println(a.Text)
}

// **`main()` is mandatory in every Go program. <br/>
// It's the function initially run on program startup.
func main() {
	// **`a`** is a variable (pseudo-object) of type **`Hello`**. <br/>
	// We initialize it by feeding the **`Text`** field.
	a := Hello {
		Text: "Hello World",
	}
	// Finally, we call the struct's **`hello()`** method to display the text fed on struct/object initialization.
	a.hello()
}
