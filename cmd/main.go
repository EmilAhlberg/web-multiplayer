package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cb := func(http.ResponseWriter, *http.Request) {
		fmt.Println("hole punch")
	}
	r := mux.NewRouter()
	r.HandleFunc("/", cb)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
