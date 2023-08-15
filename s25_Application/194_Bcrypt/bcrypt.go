package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPwd := "password1232"
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPwd))
	if err != nil {
		fmt.Println("Wrong Password")
	} else {
		fmt.Println("Correct Password")
	}

}
