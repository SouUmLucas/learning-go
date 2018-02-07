package main

import (
	"github.com/SouUmLucas/structs_avancado/model"
	"fmt"
	"encoding/json"
)

func main() {
	user1 := model.User{}
	user1.Id = 1
	user1.FirstName = "Lucas"
	user1.LastName = "Santos"
	user1.Email = "lucas.santos.silva2011@gmail.com"
	user1.SetAge(21)

	user2 := model.User{}
	user2.Id = 2
	user2.FirstName = "Ellen"
	user2.LastName = "Dayane"
	user2.Email = "ellen@gmail.com"
	user2.SetAge(16)

	fmt.Printf("User 1: %+v. É adulto?: %v\n", user1, user1.IsAdult())
	fmt.Printf("User 2: %+v. É adulto?: %v\n", user2, user2.IsAdult())

	objJSON, _ := json.Marshal(user1)
	fmt.Printf("JSON format: %v", string(objJSON))
}