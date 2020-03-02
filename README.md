# Go Crash Course
This course aims to give you not only a basic understanding of Go,
but to help you develop your first Go application, test it, expose it
via HTTP, pack it into a container and automate the build-test-pack process.

## Setup your environment
First things first :)
- Go to [the Go downloads page] and follow the instructions to set it up based on your system
- Download [Docker]
- Download [Visual Studio Code]

Now that this is out of the way, let's get coding so you can familiarize yourself with
the environment and the language.

## Go basics

### Your first Go application
What kind of course would this be without the classic "Hello World!" application?  
Create a folder named `hello-go` and open Visual Studio Code on it. If you want to fancy
this up to also get started familiarizing yourself with the console/terminal just open
your favorite terminal app and type:
```sh
mkdir hello-go && code $_
```
This will create a folder named `hello-go` in the current directory and open Visual Studio Code on it.
The `$_` will get the arguments passed to the previous command, in this case `hello-go`.
So you're probably on the VS Code screen right now. Let's create a new file named `main.go`.
To create a new file you have a few options, but to make things simples since we're on the editor, let's
right click in the folder area and choose `New File`:
![visual studio code new file screen](/assets/new-file.png "visual studio code new file screen")  
Go programs are made out of packages, and the `main` package is the entrypoint of those programs.
Packages can be imported to be used by other packages and, by convention, the package name is the same as the last
element of the import path. For standard library packages, they can be like `fmt` or `net/http`. For other types
of package, the import path is usually the repository where the package is held, like `https://github.com/gorilla/mux`.
To get started with the right feet (and to be able to easily handle our dependencies), let's first make use of
`go modules`. Just open VS Code's terminal by going to `View > Terminal` type into:
```sh
go mod init github.com/<your-github-handle>/<your-go-repository-name>
```
If you don't have a `Github` account yet, i highly recommend for you to create one and replace the placeholders in the previous command with your own Github handle and a repository name.
Now you have a `go.mod` file. This file holds your module's name, the Go version used to create it and the external dependencies
it has (with their respective versions). For now it has no dependencies, but this might change soon.
Let's get back to the `main.go` file.
The first thing in a `go` file is the package name and as we said before, by convention the entrypoint to a program is the package `main`.
Write the following code to your `main.go` file:
```go
package main

import(
    "fmt"
)

func main(){
    fmt.Print("Hello World!")
}
```
When you save it, VS Code will automatically format your code following [Go's format guidelines].
Let's break this code down so you can understand what's going on. First line is the package declaration, composed
of the keyword `package` followed by the package's name. Then we have our imports. For this specific code we're
just importing the `fmt` package, which is a toolbelt used to format/output some data to the standard output.
Then we have our first function. Go function declarations are done using the keyword `func`, an optional binding, 
an optional function name, the function parameters and a function body. Our function is named `main` also by convention, and 
what it does is just print the `string` "Hello World!" to the standard output.
To run this program, go to VS Code's terminal and type:
```sh
go run main.go
```
You should see "Hello World!" in the terminal.

### Private/Public names
Go handles private/public by convention. Anything starting with a capital letter (except for package names) is public, 
while anything starting with a lower-case letter is considered private. When importing a package, only the public parts
of it are accessible.

### Variables
Variables can be created in a few ways, and they can be initialized or not. If a variable is not initialized, it takes
the type's zero-value.  
Using the var keyword to declare one or many variables (this can be done on package level or function level):
```go
var (
	index int //int zero-value is 0
	noWord string //string zero-value is ""
	didWeLearnSomething bool = true //bool zero-value is false, but we're initializing the variable with true instead
)
```
Using the short variable declaration (this can only be done inside a function):
```go
index := 3 //short variable declaration infers the type. In this case, index will be an int
noWord := ""
someWord := "bird"
```

### Pointers
Pointers hold addresses to values and this can be called "referencing a value". The zero value of a pointer is `nil`.
```go
var index *int
fmt.Println(index)  //the zero value for pointers is nil, so that's what will be printed out
i := 4              //we assign 4 to variable i
index = &i          //the & token grabs the address of a variable, in this case we're assigning the address of i to index
fmt.Println(index)  //this will print out the memory address
fmt.Println(&i)     //checking that's the same address of i
fmt.Println(*index) //the * token is used for both declaring a pointer and for dereferencing a variable, getting it's value
i = 42              //this will change index's value as well, since they point to the same memory address
fmt.Println(*index) //checking that index now holds 42
```

### Structs
A struct is a collection of fields (private or public).
```go
type Person struct {
	ID int
	Name string
	Phone string
}
```
Structs can be "instantiated" using literals, like this:
```go
person := Person{
	ID:    1,
	Name:  "Cool Person",
	Phone: "+1 111 1111-1111",
}

fmt.Println(person.Name) //struct fields can be accessed through the dot syntax
```

### Functions
A function can take zero or more arguments, and return zero or more values as well. It's a good practice that if a function can fail for some reason, it should return it's result *and* an error. Callers of the function can check if the error is nil to know that everything went as expected.
Named functions need to be declared outside of other functions, while anonymous functions can be declared inside functions.
```go
package main

import (
	"errors"
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func processSomething(thing string) (string, error) {
	if thing == "Cool Thing" {
		return "processed successfully", nil
	}

	return "", errors.New("something bad happened")
}

func main() {
	greet := func(name string) {
		fmt.Println("Hello " + name)
	}

	greet("Cool Person")
	result := add(3, 4)
	fmt.Println(result)

	processingResult, err := processSomething("Not a Cool Thing")
	if err != nil { //Go's conditionals don't need to have a parenthesis on them
		fmt.Println(err)
	} else {
		fmt.Println(processingResult)
	}
}
```

Functions can also be bound to Structs. This allows the function to have access to the function's private fields and to act on it's values.
```go
package main

import (
	"fmt"
)

type Person struct {
	ID    int
	Name  string
	Phone string
}

func (person *Person) SayIntroduction() {
	fmt.Println("Hello! My name is " + person.Name)
}

func main() {
	person := Person{
		ID:    1,
		Name:  "Cool Person",
		Phone: "+1 111 1111-1111",
	}

	person.SayIntroduction()
}
```

### Loops
Go only has one loop construct, the `for`.  
The standard for loop consists of four components: the init statement, the condition expression, the statement that'll be executed after each iteraction and the loop body.
Here's an example:
```go
package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ { //for i starting at zero, while i is less than or equal 10, increment i by one each iteraction
		fmt.Println(i)
	}
}
```
If you want a good old `while` loop, just drop the init statement and the post statement, like this:
```go
package main

import (
	"fmt"
)

func main() {
	i := 0
	for i <= 10 { //for i starting at zero, while i is less than or equal 10, increment i by one each iteraction
		fmt.Println(i)
		i++
	}
}
```
Can you guess how we would do an infinite loop? :)

### Arrays and Slices
Array is a data structure capable of storing ordered data accessible by indexes.
In Go, arrays can be declared as following:
```go
fiveWords := [5]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
}
fmt.Println(fiveWords)
```
This declares an array of five strings. Be wary that the array's can't be resized, but Go offers a few things to work around that.
To access the third value of the array, you could do the following:
```go
thirdWord := fiveWords[2]
fmt.Println(thirdWord)
```
This grabs the value at index 2 and assigns it to the variable thirdWord. "Why index 2 if we wanted the third?" you might ask...that's because indexes start at zero :)

Slices on the other hand, reference underlying arrays. This means that they can be resized (by pointing to a bigger array).
If we wanted just the first two words of our fiveWords array, we could slice it like this:
```go
firstTwo := fiveWords[0:2]
fmt.Println(firstTwo)
```
This notation means "slice the array starting and including index 0 and finishing and excluding index 2". So we'll only get the values on indexes 0 and 1.

If we want to create a slice from scratch (instead of creating an array then slicing it) we could just omit the length:
```go
myOddsSlice := []int{1, 3, 5, 7, 9}
fmt.Println(myOddsSlice)
```
And if we want to add items to the slice, we can use the append function:
```go
myOddsSlice = append(myOddsSlice, 11, 13)
fmt.Println(myOddsSlice)
```
The append function takes a slice and any number of arguments to add to the slice. The arguments must be of the same type as the slice values.

### Maps
Maps can be used whenever you want to pair a specific key with a value and as such, map declarations have their key and value types specified:
```go
numbers := map[string]int{}
numbers["one"] = 1
numbers["two"] = 2
numbers["three"] = 3
fmt.Println(numbers)
```
If you check the printed result you might notice that the map is not ordered like the array. Never take a map's standard ordering as given, always pick the values by the keys.

### Range
The range iterator can be used to traverse arrays, slices and maps.
When used on arrays and slices, it returns the index and value of each iteration.
When used on maps, it returns the key and value of each iteration.
```go
fiveWords := [5]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
}

for index, value  := range fiveWords {
	message := fmt.Sprintf("index: %d - value: %s", index, value) //Sprintf takes a string template and values to interpolate
	fmt.Println(message)
}

numbers := map[string]int{}
numbers["one"] = 1
numbers["two"] = 2
numbers["three"] = 3

for key, value  := range numbers {
	message := fmt.Sprintf("key: %s - value: %d", key, value)
	fmt.Println(message)
}
```

###Your fancied up first Go application
To make this first example more "real world", let's create a simple HTTP endpoint that returns our string 
as it's response.

Our `main.go` file should now look like this:
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application-json")
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write([]byte(`{"message":"Hello World!"}`))
	})

	fmt.Println("http server listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
```
Let's see what's going on here...The first thing we're doing is creating an HTTP handler by binding a path named `/hello`
to a function. The HTTP handler functions always receive two parameters: an http.ResponseWriter to write our response
back to whomever requested it and an *http.Request to be able to understand what type of request we're dealing with. For this 
basic example, we're not using the *http.Request for anything. The function body is first writing 
`Content-Type: application/json` to the response's headers, then writing out the header with a `200` HTTP status code.
When you call `WriteHeader` function, you can no longer add anything to the headers. Keep that in mind. Then we're writing out
our response body. The `Write` function receives an `[]byte`. What we're doing is handwriting a `JSON` string and converting
it to `[]byte` to be used in the `Write` function.
Outside the HTTP handler function, we're just printing out a message so we're aware that something happened and then we start
our HTTP server, which will be listening for requests on port `:8080`. `http.ListenAndServe` accepts an address and an 
HTTP Handler as it's parameters. The address is `:8080`, but why are we passing `nil` as the handler?  
`http.HandleFunc` binds the handler to the default HTTP handler and when `http.ListenAndServe` doesn't receive any handlers,
it uses this default HTTP handler. That's why everything works as expected :) (or it should...let's test it out!)
On your VS Code's terminal, run the application:
```sh
go run main.go
```
You should see our message saying that the server is running, and the prompt should be in "busy mode".
To test this out, open your browser and type `localhost:8080/hello`.
You should see our message in the response sent to the browser!
To stop the application, go back to the terminal and hold the `control` key and press `c`. This should return the
terminal to the "idle mode".

[the Go downloads page]: https://golang.org/dl/
[Docker]: https://www.docker.com/products/docker-desktop
[Visual Studio Code]: https://code.visualstudio.com/Download
[Go's format guidelines]: https://golang.org/cmd/gofmt/