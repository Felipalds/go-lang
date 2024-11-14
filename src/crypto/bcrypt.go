package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)


func main () {

    // let's define our password
    password := "123456"

    // here we call the function bcrypt from golang
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
       fmt.Println("Error generating hash!", err) 
       return
    }

    fmt.Println("Bcrypt:", string(hash))

}
