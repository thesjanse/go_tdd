package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Hello(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!", name)
}

func MyHelloHandler(w http.ResponseWriter, r *http.Request) {
	Hello(w, "Lucius")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyHelloHandler)))
}
