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