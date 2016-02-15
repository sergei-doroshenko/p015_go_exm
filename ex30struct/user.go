package main

import (
	"fmt"
)

type User struct {
	Name string
}

func changeExist(u *User) {
	u.Name = "Raul"
}

func notChange(u User) {
	u.Name = "Changed"
}

func getPointer(u User) *User {
	return &User{"Anton"}
}

func getUser(name string) User {
	return User{name}
}

func main() {
	user := User{"Huan"}
	fmt.Println("user.Name: ", user.Name) // Huan

	changeExist(&user)                    // change name in existent user
	fmt.Println("user.Name: ", user.Name) // Raul

	notChange(user) //
	fmt.Println("user.Name: ", user.Name)

	u := getPointer(user)
	fmt.Println("user.Name: ", user.Name)
	fmt.Println("u.Name: ", u.Name)

	fmt.Println("From getUser: ", getUser("Sergio").Name)
}
