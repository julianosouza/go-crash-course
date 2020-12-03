package main

import (
	"fmt"
)

// to run this program, `cd` to this directory and run `go run main.go`
func main() {
	// you can initialize variables by using the keyword `var`,
	// the name of the variable and then it's type
	var a int
	var b string

	// then you can assign values to the already created variables using the `=` sign
	a = 5
	b = "five"

	// you can also initialize variables by using the shorthand form,
	// which sets the type of the variable as the type being assigned to it
	msg := ""

	msg = fmt.Sprintf("a = %d and b = %s", a, b)

	fmt.Println(msg)
}
