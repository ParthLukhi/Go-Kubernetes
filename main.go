package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/ParthLukhi/Go-Kubernetes/controllers"
	"github.com/ParthLukhi/Go-Kubernetes/pkg/cluster"
)

// func getConfigPath() string {
// 	// Provide the absolute path to the Kubernetes configuration file
// 	return "/Users/parthlukhi/.kube"
// }

func main() {

	r := mux.NewRouter()
	controllers.RegisterMovieRoutes(r)

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

	kubeconfig := "/Users/parthlukhi/.minikube/config"
	os.Setenv("KUBECONFIG", kubeconfig)

	// r := mux.NewRouter()

	// controllers.RegisterMovieRoutes(r)

	// http.Handle("/", r)
	// fmt.Println("Starting server on :8000...")
	// go func() {
	// 	log.Fatal(http.ListenAndServe(":8000", nil))
	// }()

	// Initialize Cluster Scanner
	scanner, err := cluster.NewClusterScanner(kubeconfig)
	if err != nil {
		log.Fatalf("Failed to create cluster scanner: %v", err)
	}

	// Scan the cluster
	if err := scanner.ScanNodes(context.Background()); err != nil {
		log.Fatalf("Failed to scan nodes: %v", err)
	}

	// r := mux.NewRouter()

	// controllers.RegisterMovieRoutes(r)

	// fmt.Printf("Starting server at port 8000\n")
	// log.Fatal(http.ListenAndServe(":8000", r))
}
