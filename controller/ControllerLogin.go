package controller

import "net/http"
import "fmt"

func RouteManagerAccount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		RequestGetInfoAccount(w, r)
	case "POST":
		RequestUpdateAccount(w, r)
	case "DELETE":
		RequestDeleteAccount(w, r)
	default:
		fmt.Println("Error en petición <ACCOUNT>")
	}
}

func RouteManagerLog(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		RequestLog(w, r)
	default:
		fmt.Println("Error en petición <LOGIN>")
	}
}

func RouteManagerRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		RequestRegister(w, r)
	default:
		fmt.Println("Error")
	}
}

func RequestLog(w http.ResponseWriter, r *http.Request)            {}
func RequestRegister(w http.ResponseWriter, r *http.Request)       {}
func RequestDeleteAccount(w http.ResponseWriter, r *http.Request)  {}
func RequestUpdateAccount(w http.ResponseWriter, r *http.Request)  {}
func RequestGetInfoAccount(w http.ResponseWriter, r *http.Request) {}
