package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, Spinnaker")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
