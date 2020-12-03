package main

import (
	"fmt"
)

func main() {
	// when dealing with variables, you must always consider if you're
	// dealing with a reference type or a value type
	// otherwise you might get some unwanted results
	a := 5
	b := a  //assigned the value of a to b
	c := &a //assigned the `address` of a to c

	// the type of c is not int, but instead *int
	// which means it's a pointer to a int value
	// we can check that by printing it, for instance
	// instead of outputing 5, it will output it's memory address
	fmt.Println(fmt.Sprintf("the value of c is %v", c))

	// we can assign a function to a variable and reuse it later
	printAll := func() {
		// when we want to get a value of a pointer, we need to dereference it
		fmt.Println(fmt.Sprintf("a is %d, b is %d, c internal value is %d", a, b, *c))
	}

	printAll()

	// if we update a, we update all of it's references
	// every variable which owns it's address
	a = 8
	// notice that by updating a, c also got updated, because it had the same reference as a.
	// b stayed the same because it only had the value of a
	printAll()
}
