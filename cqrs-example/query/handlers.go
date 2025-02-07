package query

import "cqrs-example/command"

func GetUser(id string) (command.User, bool) {
	user, exists := command.UserStore[id]
	return user, exists
}
