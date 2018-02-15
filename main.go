package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/andreslab/prj_api_crop/controller"
)

func CheckState(w http.ResponseWriter, r *http.Request) {
	fmt.Println("state 200")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", CheckState)
	mux.HandleFunc("/crop/scan", controller.ManagerRouterScan)
	mux.HandleFunc("/crop/user", controller.ManagerRouterUser)
	mux.HandleFunc("/crop/storage", controller.ManagerRouterStorage)
	mux.HandleFunc("/crop/weather", controller.ManagerRouterWeather)
	mux.HandleFunc("/crop/profile", controller.ManagerRouterProfile)
	mux.HandleFunc("/crop/pests", controller.ManagerRouterPests)

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
