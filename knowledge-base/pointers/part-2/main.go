package main

import "fmt"

// when
// 1. we want to modify the state
// 2. we want to avoid copying large structs (to optimize memory) (pointer = 8 bytes)
// 3. we want to share the state across multiple functions

type User struct {
	username string
	age      int
	email    string
	file     []byte
}

func (u User) Email() string { // <- file will be copied as well, which is inefficient "when.2"
	return u.email
}

func (u *User) toString() string { // <- file will not be copied, only the pointer to the struct will be passed "when.2"
	return fmt.Sprintf("Username: %s, Age: %d, Email: %s", u.username, u.age, u.email)
}

func (u *User) SetEmail(email string) { // <- "when.1"
	u.email = email
}

func main() {
	user := User{
		username: "john_doe",
		age:      30,
		email:    "john.doe@example.com",
	}
	fmt.Println(user.Email())
	user.SetEmail("new.email@example.com")
	fmt.Println(user.Email())
	userData := user.toString()
	fmt.Println(userData)
}
