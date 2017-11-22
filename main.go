package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, friend!")
		fmt.Printf("%s %s %d\n", r.Method, r.URL.Path, http.StatusOK)
	})

	fmt.Println("Started.")
	fmt.Println(http.ListenAndServe(":80", nil))
}
