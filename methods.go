package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) UserDetails() {
	fmt.Println("username is : ", u)

}

func met() {

	us := User{"asif", 26}
	us.UserDetails()

}
