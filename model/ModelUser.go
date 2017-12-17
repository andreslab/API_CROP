package model

import "fmt"

func Imprimi() {
	fmt.Printf("hola")
}

type UserModel struct {
	id      string `json:"id"`
	name    string `json:"name"`
	email   string `json:"email"`
	pass    string `json:"pass"`
	created int    `json:"created"`
	scan    int    `json:"scan"`
}
