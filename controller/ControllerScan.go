package controller

import (
	"fmt"
	"net/http"
)

func ManagerRouterScan(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetScan(w, r)
	case "POST":
		RequestPostScan(w, r)
	case "UPDATE":
		RequestUpdateScan(w, r)
	case "DELETE":
		RequestDeleteScan(w, r)
	default:
	}
}

func RequestGetScan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Get")
}

func RequestPostScan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Post")
}

func RequestUpdateScan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Update")
}

func RequestDeleteScan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Delete")
}
