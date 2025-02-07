package command

import "sync"

var UserStore = make(map[string]User)
var mu sync.Mutex

func CreateUser(id, name, email string) {
	mu.Lock()
	defer mu.Unlock()
	UserStore[id] = User{id, name, email}
}
