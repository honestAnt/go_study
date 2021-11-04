package main

import (
	"fmt"
	"log"
	"net/http"
)

func BuildServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello here... %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", BuildServer)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
