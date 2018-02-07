package main

import (
    "fmt"
	"github.com/SouUmLucas/learning-go/maps/model"
)

var users map[int]model.User

func main() {
	users = make(map[int]model.User)

	makeUser(1, "Lucas Santos", "lucas@giter.com.br")
	makeUser(2, "Matheus", "mt@gmail.com")
	makeUser(3, "Fabio", "fabio@gmail.com")

	mapInfo(users)
	listMap()

	user, found := users[2]
	if found {
		delete(users, user.Id)
	}

	mapInfo(users)
	listMap()
}

func makeUser(id int, name string, email string) model.User {
	u := model.User{Id: id, Name: name, Email: email}
	addUserToMap(u)
	return u
}

func addUserToMap(user model.User) {
	users[user.Id] = user
}

func mapInfo(m map[int]model.User) {
	fmt.Printf("Map length %d\n", len(m))
}

func listMap() {
	for key, user := range users {
		fmt.Printf("Key: %d - %+v\n", key, user)
	}
}