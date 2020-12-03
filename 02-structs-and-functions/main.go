package main

import (
	"fmt"
)

func main() {
	a := 3
	b := 9
	result := Sum(a, b)
	fmt.Println(fmt.Sprintf("the result of summing %d and %d is %d", a, b, result))

	// the struct initialization with assignment by reference
	johnDoe := &person{
		Name: "John Doe",
		Age:  21,
	}

	fmt.Println(johnDoe.SayHello())
	fmt.Println("--John Doe got older")
	johnDoe.GetOld()
	fmt.Println(johnDoe.SayHello())
}

// Sum is a function created to represent the mathematical operation of Sum
// functions consist of the keyword `func`, the name of the function (in this case Sum)
// the input parameters (in this case a and b) and the return type (in this case int)
// Go functions allow multiple return types on a single function
// it's even a good practice for every function to return the type `error` as it's last return type
// this allows for the caller to know if the function executed without errors or not by checking error != nil
// for this example we're not returning error, however
func Sum(a int, b int) int {
	return a + b
}

// Person is a struct created to allow us to represent something more complex than a
// number or text. Something that can even have some behaviors.
// Structs can have multiple fields (in this case Name and Age) with distinct types
// They can also have bounded functions (which in other languages we usually call methods)
// Note that capitalization on Go is important. Private and Public on Go is controlled by convention.
// If something starts with a capital letter, it's public, which means it can be used outside of
// the package it was declared in, otherwise it's private to the package it was declared in.
type person struct {
	Name string
	Age  int
}

// GetOld is bound to any reference of person
// This is basically how you create a method attached to a struct
// You can't use this function without a reference of person
func (p *person) GetOld() {
	p.Age++
}

// SayHello is also bound to person
func (p *person) SayHello() string {
	msg := fmt.Sprintf("Hi! My name is %s and i'm %d years old.", p.Name, p.Age)
	return msg
}
