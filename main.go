package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", CheckState)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	server.ListenAndServe()
}

func CheckState(w http.ResponseWriter, r *http.Request) {
	fmt.Println("state 200")
}
