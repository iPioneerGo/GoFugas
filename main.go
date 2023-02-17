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

var usersMap = map[int]user{}

func (u user) String() string {
	return fmt.Sprintf(""+
		"{\n"+
		"  name: %v,\n"+
		"  age: %d,\n"+
		"  email: %v "+
		"\n}", u.name, u.age, u.email)
}

func searchAndPrint(users []user, name string) {
	serviceMap := convertUsersSliceToMap(users, usersMap)
	for _, user := range serviceMap {
		if strings.EqualFold(user.name, name) {
			fmt.Println(user)
		}
	}
}

func convertUsersSliceToMap(users []user, usersMap map[int]user) map[int]user {
	for i := range users {
		usersMap[i] = users[i]
	}
	return usersMap
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
