package main

import (
	"fmt"
)

func main() {
	line := []string{"Ana", "Billy", "Joe"}

	// if and else are the basic keywords used to control flow
	// if evaluates an expression. If the expression is true, it's following code block is ran.
	// otherwise, the else block is ran. Keep in mind that the else block is optional.
	// You can have if without else
	if len(line)%2 != 0 {
		fmt.Println("it's an odd line")
	} else {
		fmt.Println("it's an even line")
	}

	// another way of expressing certain multi-condition evaluations is by using `case`
	// if you don't want other cases to be evaluated after a match, use the keyword break
	switch len(line) {
	case 1:
		fmt.Println("there's one person in line")
		break
	case 2:
		fmt.Println("there're two persons in line")
		break
	case 3:
		fmt.Println("there're three persons in line")
		break
	default:
		fmt.Println("none of the cases were a match")
	}

	// simple for loop
	// we can use the builtin func `len` to get length of the `line`
	for i := 0; i < len(line); i++ {
		fmt.Println(fmt.Sprintf("current person in line by using a regular for loop is %s", line[i]))
	}

	// another, way of doing that is by using the range operator
	for i, value := range line {
		fmt.Println(fmt.Sprintf("range current index is %d and current value is %s", i, value))
	}

	// if you want to do a while loop, you can ommit some of the arguments
	i := 0
	for i < 5 {
		fmt.Println(fmt.Sprintf("current index using for loop with only the condition is %d", i))
		i++
	}

	// if you want, you can ommit everything
	// this would be the equivalent of the `do while` construct of some languages
	// just be sure to have some way of breaking out of the for loop
	// otherwise it'll run until it explodes :)
	i = 0
	for {
		fmt.Println(fmt.Sprintf("current index using for loop without any condition is %d", i))
		i++
		if i == 5 {
			// this will break out of the loop when i equals 5
			break
		}
	}
}
