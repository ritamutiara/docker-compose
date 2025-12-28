package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "local"
	}

	fmt.Println("Run app in port : " + port)
	fmt.Println("Mode: " + mode)

	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/health", HealthCheck)

	http.ListenAndServe(":"+port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	mode := os.Getenv("MODE")
	response := "Hello " + mode
	fmt.Fprintf(w, response)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}