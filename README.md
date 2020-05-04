# Go Crash Course
This course aims to give you not only a basic understanding of Go,
but to help you develop your first Go application, test it, expose it
via HTTP, pack it into a container and automate the build-test-pack process.

## Setup your environment
First things first :)
- Go to [the Go downloads page] and follow the instructions to set it up based on your system
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

### Approaching the real world scenario
For the sake of brevity, we won't dive down on what's HTTP or JSON but feel free to do a quick research on those before continuing if you want :)

To make our Hello World more similar to what most developers would do on a "real world" scenario, 
let's create a simple HTTP endpoint that returns our string 
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
When you call `WriteHeader` function, you can no longer add anything to the headers, keep that in mind. Then we're writing out
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

### A simple example application: a home devices controller
Automating your home devices is quite trendy nowadays, so a cool example application for us to learn some more is an API to control our house appliances. The basic requirement is for us to deliver an HTTP service that is able to say our current devices status and change those status whenever we want. Let's start it simple, with controlling our lightbulbs.

First of all, let's create a simple map to hold the state of the lightbulbs in our house and assign a few standard lightbulbs to it.
```go
package main

import (
	"fmt"
	"net/http"
)

var (
	lightbulbs = make(map[string]bool)
)

func main() {
	lightbulbs["livingroom"] = false
	lightbulbs["kitchen"] = false

	http.HandleFunc("/hello", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application-json")
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write([]byte(`{"message":"Hello World!"}`))
	})

	fmt.Println("http server listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
```

If you recall from our previous blog post, maps can be used to index values by key. In this case, we're indexing lightbulb states (on or off) by room name and creating each lightbulb turned off by default.
Now we want an endpoint to list these states. Let's create a simple endpoint that will output the current state of our lightbulbs grid.

```go
http.HandleFunc("/lightbulbs", func(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application-json")
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(lightbulbs)
})
```

As on our Hello World example, this is a simple endpoint that binds an http handle function to the path `/lightbulbs`, meaning that whenever a request to that path is made, the function will run. There's a new element to this handle function however. Since we don't want to hand roll our JSON structure, we're using a JSON encoder to write the JSON representation of our map to the ResponseWriter. If we run our code and hit this path on the browser, we should be able to see both lightbulbs having the value of `false` as they're turned off right now.

Now we need a way to switch our lightbulbs, right? Let's create a new endpoint to also deal with that behavior.

```go
http.HandleFunc("/lightbulbs/switch", func(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application-json")

	name := request.URL.Query().Get("name")
	if name == "" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte(`{"message":"a lightbulb name should be provided as the value of a 'name' querystring parameter"}`))
		return
	}

	if _, keyExists := lightbulbs[name]; !keyExists {
		responseWriter.WriteHeader(http.StatusNotFound)
		responseWriter.Write([]byte(`{"message":"a lightbulb with the provided name doesn't exist"}`))
		return
	}

	lightbulbs[name] = !lightbulbs[name]

	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(lightbulbs)
})
```

Step by step, we set the content type of our future response, then we get a value from our querystring (key-value pairs appended to the URL after a `?`) with the key "name". If the querystring value was not present or was empty, we set our response status code to BadRequest to indicate that the request didn't follow the rules we expected and set the response to a hand rolled JSON containing a message explaining what happened, then we use the keyword `return` to end the execution of the http handler right there, as there's no point in continuing with the request processing if it doesn't contain the information needed to continue with it.

If the name does exist, we try to get a value from our map using that name as the key. If a value with that key doesn't exist, we set the response status code to NotFound and the response to a hand rolled JSON containing a message explaining that no lightbulb was found with that name.

If there is a key on the map with the given name, we set it's value to the opposite of it by using the `! ` operator, meaning that if the value is `false`, we assign `true`. Then we do the same as with the `status` endpoint, by setting the response status code to OK and writing the JSON representation of our map to the response.

Now we need to create an endpoint to add new lightbulbs, right?

```go
http.HandleFunc("/lightbulbs/create", func(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application-json")

	name := request.URL.Query().Get("name")
	if name == "" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte(`{"message":"a lightbulb name should be provided as the value of a 'name' querystring parameter"}`))
		return
	}

	if _, keyExists := lightbulbs[name]; keyExists {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte(`{"message":"a lightbulb with the provided name already exists"}`))
		return
	}

	lightbulbs[name] = false

	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(lightbulbs)
})
```

It does basically the same as our previous endpoint, as it checks for the presence of a `name` querystring parameter, but when checking if the key exists on our map or not, for an erroring scenario we're interested on the case that the key already exists on our map, because we're not supposed to create something that already exists. If the key is already present, our response is a BadRequest with a message saying what's the reason for the BadRequest response.

In case everything goes as planned, we simply create a new entry on our map for a turned off lightbulb with the given name and respond with OK and the current content of our map.

Since we're creating new lightbulbs, we should be able to delete them as well, right? Let's create an endpoint for that.

```go
http.HandleFunc("/lightbulbs/delete", func(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application-json")

	name := request.URL.Query().Get("name")
	if name == "" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte(`{"message":"a lightbulb name should be provided as the value of a 'name' querystring parameter"}`))
		return
	}

	if _, keyExists := lightbulbs[name]; !keyExists {
		responseWriter.WriteHeader(http.StatusNotFound)
		responseWriter.Write([]byte(`{"message":"a lightbulb with the provided name doesn't exists"}`))
		return
	}

	delete(lightbulbs, name)

	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(lightbulbs)
})
```

The difference on this endpoint is basically the second validation and the success case, as we can't delete something that doesn't exist. If everything goes as planned, we call `delete` passing the map and the key we want to delete.

For future uses, let's change our hello handler to be a `healthcheck` handler. We can use it to make sure our service is up and running in a reliable way.

```go
http.HandleFunc("/healthcheck", func(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application-json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte(`{"message":"service is up and running"}`))
})
```

### Testing

On future posts we're going to change this application to improve it's readability and testability, and this act of improving your codebase without changing behavior is called `refactor`, but in order to do that we need to make sure that our future changes don't break our current behavior. A good way to do that is to implement tests that can be run whenever we want to ensure everything is working as expected.

Go tests are functions in the form of:
```go
func TestXxx(t *testing.T){}
```

By convention, the tool `go test` treats functions following that template as tests.
Let's write our first test for our healthcheck endpoint. To do this, first create a new file named `healthcheck_test.go`, then use this as the file's content:

```go
package main

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	go main()

	response, err := http.Get("http://localhost:8080/healthcheck")
	if err != nil {
		t.Errorf("expected no errors, but got %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("expected 200 statuscode, but got %v", response.StatusCode)
	}

	responseBody := make(map[string]interface{})
	json.NewDecoder(response.Body).Decode(&responseBody)
	response.Body.Close()

	if responseBody["message"] != "service is up and running" {
		t.Errorf(`expected message to be "service is up and running", but got %v`, responseBody["message"])
	}

	os.Interrupt.Signal()
}
```

This is a pretty naive approach to testing. I'm keeping it simple for now, but we'll be refactoring this as well in the future.

We're using the same package name to keep things simple for now. We've created our first test function, named `TestHealthCheck` to follow Go's naming convention for tests. The test itself is pretty straightforward, but there are a bunch of new things, so let's walk through it slowly. First is a new keyword, `go`. Remember that when we run our application, the terminal sits in `busy mode` and we can't do anything with it unless we terminate the program? It's the same here. If we just run the main function, our test would be stuck, waiting until the test itself times out after 30 seconds. The `go` keyword tells the program to run the provided function on a separate `goroutine`. We'll be talking about goroutines further down the course, but for now, let's say it runs the provided function on a different `workspace` than the currently running function, in a way to not block the execution of it. This way we can run our main function and still collect our test results. After running the main function, we're using the package http to do a GET http call to our service, on the `healthcheck` endpoint. The result of the `http.Get` function is an http response and an error, so the first thing we do is check for the error, then check if the response status code is what we expect and finally check if the response body is what we expect. Whenever a result is not what we're expecting, we're calling `t.Errorf` to fail the test with a friendly message, telling us what went wrong.
To write the JSON representation of a structure to the ResponseWriter, we used `json.NewEncoder().Encode()`. To reverse this process and grab a structure from a JSON encoded response, we use `json.NewDecoder().Decode()`. After we've ran all the checks, it's time to fire an `Interrupt` signal to stop our program, much like doing a `control+c` from our terminal.

To run the test, simply type `go test ./...` in the terminal. The test is going to run and the output is going to be that all the tests passed. In order to check that our test is really validating something, feel free to change the response status code from our healthcheck function to something else than an `http.StatusOK` and run the tests again. We'll notice that now the test fails, as it's expecting a status OK, but got something else.

[the Go downloads page]: https://golang.org/dl/
[Docker]: https://www.docker.com/products/docker-desktop
[Visual Studio Code]: https://code.visualstudio.com/Download
[Go's format guidelines]: https://golang.org/cmd/gofmt/