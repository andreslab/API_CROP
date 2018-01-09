package controller

import (
	"fmt"
	"net/http"
)

func ManagerRouterResult(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetResult(w, r)
	case "POST":
		RequestPostResult(w, r)
	case "UPDATE":
		RequestUpdateResult(w, r)
	case "DELETE":
		RequestDeleteResult(w, r)
	default:
	}
}

func RequestGetResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Get")
}

func RequestPostResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Post")
}

func RequestUpdateResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Update")
}

func RequestDeleteResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Delete")
}
