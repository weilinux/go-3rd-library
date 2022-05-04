package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	ID string `validate:"omitempty,uuid"`
	FirstName string `validate:"required"`
	LastName string `validate:"required"`
	Username string `validate:"required,email"`
	Password string `validate:"required,gte=10"`
	Type string `validate:"isdefault"` //the default value of this type
}

//more demo https://github.com/Jambo-Git/validator-demo
//there are many different validator that found on the github.com
func main() {
	user := User{
		Username:  "weilin@sina.com",
		Password:  "1234567890",
		Type: "",
	}

	validate := validator.New()
	//normal check
	//err := validate.Struct(user)

	//StructExcept check
	err := validate.StructExcept(user, "FirstName", "LastName")
	if err != nil {
		fmt.Println(err.Error())
	}

	//validate for a variable name
	password := "13243"
	err = validate.Var(password, "required,gte=10")
	if err != nil {
		fmt.Println(err.Error())
	}


}
