package controller

import (
	"fmt"
	"net/http"
)

func ManagerRouterInfo(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetInfo(w, r)
	case "POST":
		RequestPostInfo(w, r)
	case "UPDATE":
		RequestUpdateInfo(w, r)
	case "DELETE":
		RequestDeleteInfo(w, r)
	default:
	}
}

func RequestGetInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Get")
}

func RequestPostInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Post")
}

func RequestUpdateInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Update")
}

func RequestDeleteInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Delete")
}
