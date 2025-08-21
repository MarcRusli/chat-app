package main

import (
	"fmt"
	"net/http"
)

func setRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
}

func main() {
	setRoutes()
	http.ListenAndServe(":8080", nil)
}
