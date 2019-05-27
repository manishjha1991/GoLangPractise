package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Test")
}
func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":9000", nil)
}
