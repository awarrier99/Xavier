package user

import "fmt"

type User struct {
    first_name, last_name, nick_name, DOB string
    age, PIN int
}

func NewUser(first_name, last_name, nick_name, DOB string, age, PIN int) *User {
    return &User{first_name, last_name, nick_name, DOB, age, PIN}
}

func (u User) String() string {
    return fmt.Sprintf("%#v", u)
}
