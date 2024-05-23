package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ParthLukhi/Go-Kubernetes/controllers"
)

func main() {

	r := mux.NewRouter()

	controllers.RegisterMovieRoutes(r)

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
