
# Basic REST API Setup in Golang

## Prerequisites

- Go is installed on your system. If not, download and install it from the official Go website: https://golang.org/dl/

## Step 1: Create a new project directory

Set up a new directory for your Golang project.

## Step 2: Initialize Go module

Go modules help manage dependencies in your project. Run the following command inside your project directory to initialize the module:

```bash
go mod init your_project_name
```
## Step 3: Install necessary libraries

We'll use the Gorilla Mux package for routing. Install it using the following command:

```bash
go get -u github.com/gorilla/mux
```
## Step 4: Create the main.go file

This file will serve as the entry point to your REST API application.

``` bash
package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    // Initialize the router
    router := mux.NewRouter()

    // Define your API routes
    router.HandleFunc("/api/hello", helloHandler).Methods("GET")

    // Start the server
    log.Println("Server started on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Your API logic goes here
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, World!"))
}

```

## Step 5: Run the API
Save the main.go file and run your API with the following command:

```bash
go run main.go
```

Your REST API should now be running at http://localhost:8080/api/hello. You can access it using a web browser or tools like curl or Postman.

This is just a basic setup to get you started with a REST API in Golang. As your project grows, you might want to separate your code into different files, add a database, authentication, middleware, and other features depending on your requirements. But this simple setup should give you a good starting point.
