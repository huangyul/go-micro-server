package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 密码加密
	password := "password"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	// 密码解密
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hash))
}
