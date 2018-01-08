package controller

import (
	"fmt"
	"net/http"
)

func ManagerRouterAlarm(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetAlarm(w, r)
	case "POST":
		RequestPostAlarm(w, r)
	case "UPDATE":
		RequestUpdateAlarm(w, r)
	case "DELETE":
		RequestDeleteAlarm(w, r)
	default:
	}
}

func RequestGetAlarm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Get")
}

func RequestPostAlarm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Post")
}

func RequestUpdateAlarm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Update")
}

func RequestDeleteAlarm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request Delete")
}
