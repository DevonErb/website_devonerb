// hellow world works with port 5050 
// probably because of my golive on vs code
// couldnt get 80 or 8080 working
//------------------
//setup
//  go version - check what version running
//  go mod init website
//  build: go build
//  run: ./website 
//  build-run: go run main.go
//killing website 
//  pgrep helloworld
//  kill -9 pid from pgrep
//------------------

package main

import (
    "net/http"
    "log"
)

func main() {


    //This servers the static folder index.html style.css and devonerb.js
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    // Log the server start
    log.Println("Starting server on :5050")
    
    // Start the server
    err := http.ListenAndServe(":5050", nil)
    if err != nil {
        log.Fatal(err)
    }
}