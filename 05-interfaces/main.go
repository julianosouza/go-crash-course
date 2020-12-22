package main

import (
	"fmt"
)

func main() {
	catto := &cat{
		Name: "Catto",
	}

	doggo := &dog{
		Name: "Doggo",
	}

	// we can assign cat and dog to the speaker type, because they implement the interface
	animals := []speaker{catto, doggo}

	// since we're not interested in the index of the loop, we can ignore it by using the `_`
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

// interfaces describe certain expected behaviors
// this one says that a speaker can speak
type speaker interface {
	Speak() string
}

type cat struct {
	Name string
}

// interface implementation in Go works by just having the exact same function signature as an interface
// in this case, since cat has a function named Speak, which returns string, it matches interface Speaker
// so cat is a speaker
func (c *cat) Speak() string {
	return fmt.Sprintf("cat %s says meooooow", c.Name)
}

type dog struct {
	Name string
}

// dog is also a speaker by having a func Speak with the exact same args (none) and return type (string)
// as the func Speak of the interface speaker
func (c *dog) Speak() string {
	return fmt.Sprintf("dog %s says bark! bark!", c.Name)
}
