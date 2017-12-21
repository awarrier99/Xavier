package main

import (
    "fmt"
    "github.com/awarrier99/Xavier/user"
)

func main() {
    u := user.NewUser("Ashvin", "Warrier", "Ash", "5/18/1999", 18, 9277)
    fmt.Println(u)

    u3 := new(user.User)
    fmt.Println(u3)
}
