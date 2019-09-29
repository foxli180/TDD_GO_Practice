package main

import (
	"awesomeProject/di"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
