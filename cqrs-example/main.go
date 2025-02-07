package main

import (
	"cqrs-example/command"
	"cqrs-example/query"
	"fmt"
)

func main() {
	command.CreateUser("1", "John Doe", "JonDoe@example.com")

	user, exists := query.GetUser("1")

	if exists {
		fmt.Println("User found:", user.Name, "-", user.Email)
	} else {
		fmt.Println("User not found")
	}
}
