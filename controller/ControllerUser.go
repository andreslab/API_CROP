package controller

import (
	"fmt"
	"net/http"
)

func ManagerRouterUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetUser(w, r)
	case "POST":
		RequestPostUser(w, r)
	case "UPDATE":
		RequestUpdateUser(w, r)
	case "DELETE":
		RequestDeleteUser(w, r)
	default:
	}
}

func RequestGetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Get")
}

func RequestPostUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Post")
}

func RequestUpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Update")
}

func RequestDeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Delete")
}
