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

As an excercise, try to test our other endpoints!

On the next blog post i'll cover testing the other functions for every scenario they're presenting and we'll get started with our refactor. Until next time! :)