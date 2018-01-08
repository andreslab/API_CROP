package controller

import (
	"fmt"
	"net/http"
)

func ManagerRouterNote(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetNote(w, r)
	case "POST":
		RequestPostNote(w, r)
	case "UPDATE":
		RequestUpdateNote(w, r)
	case "DELETE":
		RequestDeleteNote(w, r)
	default:
	}
}

func RequestGetNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Get")
}

func RequestPostNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Post")
}

func RequestUpdateNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Update")
}

func RequestDeleteNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Delete")
}
