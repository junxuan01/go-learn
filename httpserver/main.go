package main

import (
	"fmt"
	"net/http"
)

func main() {
	a := 3
	fmt.Println(a)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	http.ListenAndServe(":8090", nil)
}
