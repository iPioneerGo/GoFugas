package main

import (
	"fmt"
	"strings"
)

type user struct {
	name  string
	age   uint8
	email string
}

func (u user) String() string {
	return fmt.Sprintf(""+
		"{\n"+
		"  name: %v,\n"+
		"  age: %d,\n"+
		"  email: %v "+
		"\n}", u.name, u.age, u.email)
}

func searchAndPrint(users []user, name string) {
	for _, user := range users {
		if strings.ToLower(user.name) == strings.ToLower(name) {
			fmt.Println(user)
		}
	}
}

func main() {

	users := []user{
		{
			name:  "Mike",
			age:   32,
			email: "mike@gmail.com",
		},
		{
			name:  "John",
			age:   54,
			email: "john@gmail.com",
		},
		{
			name:  "Abobus",
			age:   2,
			email: "abobus@gmail.com",
		},
	}
	searchAndPrint(users, "Mike")
	searchAndPrint(users, "mike")
	searchAndPrint(users, "John")
	searchAndPrint(users, "john")
	searchAndPrint(users, "Abobus")
	searchAndPrint(users, "abobus")
}
