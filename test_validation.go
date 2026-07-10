package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Test struct {
	ImageURL string `validate:"omitempty,url"`
}

func main() {
	v := validator.New()
	t := Test{ImageURL: ""}
	err := v.Struct(t)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No Error")
	}
}
