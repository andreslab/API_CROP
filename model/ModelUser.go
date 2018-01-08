package model

import (
	"fmt"
)

func (c UserModel) Imprimi() {
	fmt.Printf("hola")
}

type UserModel struct {
	ID       int64
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Pass     string `json:"pass"`
	Created  string `json:"created"`
}
