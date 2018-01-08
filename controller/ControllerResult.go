package controller

import (
	"fmt"
	"net/http"
)

func ManagerRouterresult(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetresult(w, r)
	case "POST":
		RequestPostresult(w, r)
	case "UPDATE":
		RequestUpdateresult(w, r)
	case "DELETE":
		RequestDeleteresult(w, r)
	default:
	}
}

func RequestGetresult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Get")
}

func RequestPostresult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Post")
}

func RequestUpdateresult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Update")
}

func RequestDeleteresult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Delete")
}
