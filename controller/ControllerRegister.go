package controller

import (
	"fmt"
	"net/http"
)

func ManagerRouterregister(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetregister(w, r)
	case "POST":
		RequestPostregister(w, r)
	case "UPDATE":
		RequestUpdateregister(w, r)
	case "DELETE":
		RequestDeleteregister(w, r)
	default:
	}
}

func RequestGetregister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Get")
}

func RequestPostregister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Post")
}

func RequestUpdateregister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Update")
}

func RequestDeleteregister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Delete")
}
