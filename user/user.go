package user

import "fmt"

type User struct {
    first_name, last_name, nick_name, dob string
    age int
}

func NewUser(first_name, last_name, nick_name, dob string, age int) *User {
    return &User{first_name, last_name, nick_name, dob, age}
}

func (u *User) String() string {
    return fmt.Sprintf("%#v", u)
}

func (u *User) GetFirstName() string {
    return u->first_name
}

func (u *User) GetLastName() string {
    return u->last_name;
}
func (u *User) GetNickName() string {
    return u->nick_name;
}

func (u *User) GetDOB() string {
    return u->DOB;
}

func (u *User) GetAge() int {
    return u->age;
}
