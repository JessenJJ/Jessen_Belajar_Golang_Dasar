package basic

import "fmt"

type User struct {
	Firstname string
	Lastname  string
	Age       int
}

func NewUser(Firstname, Lastname string, Age int) *User {
	return &User{Firstname, Lastname, Age}

}

func (u User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.Firstname, u.Lastname)

}
