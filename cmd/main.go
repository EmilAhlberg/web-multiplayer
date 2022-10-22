package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	cb := func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("hole punch", r.Method)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", cb)
	http.Handle("/", r)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:8080"}) //os.Getenv("ORIGIN_ALLOWED")})

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headersOk, originsOk, methodsOk)(r)))
}
