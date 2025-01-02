package main

import (
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":5000", nil)

}
