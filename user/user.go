package user

import (
    "fmt"
    "time"
)

type user struct {
    first_name string
    last_name string
    nick_name string
    age int
    DOB Time
    PIN int
}

func NewUser(first_name, last_name, nick_name string, age int, DOB Time, PIN int) *user {
    return &user{first_name, last_name, nick_name, age, DOB, PIN}
}
