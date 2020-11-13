// every go program begins on package main
package main

// you can import packages that you want to use
import (
	"flag"
	"fmt"
)

// the main package must also have a main func, which is the entrypoint of the app
// to run this program, `cd` to this directory and run `go run main.go`
func main() {
	// using flags is a simple way to configure your application
	// you can set it up differently per environment, for instance
	msg := flag.String("message", "default message", "you can pass a message to be displayed")

	// just a simple print, to show the value we grabbed through flags
	// here we're using the imported package `fmt`
	fmt.Println(*msg)
}
