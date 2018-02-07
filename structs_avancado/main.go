package main

import (
	"github.com/SouUmLucas/learning-go/structs_avancado/model"
	"fmt"
	"encoding/json"
)

func main() {
	user1 := model.User{}
	user1.Id = 1
	user1.FirstName = "Lucas"
	user1.LastName = "Santos"
	user1.Email = "lucas.santos.silva2011@gmail.com"
	if err := user1.SetAge(21); err != nil {
		fmt.Println("[main] houve um error na atribuição da idade:", err.Error())
	}

	user2 := model.User{}
	user2.Id = 2
	user2.FirstName = "Ellen"
	user2.LastName = "Dayane"
	user2.Email = "ellen@gmail.com"
	if err := user2.SetAge(16); err != nil {
		fmt.Println("[main] houve um error na atribuição da idade:", err.Error())
	}

	user3 := model.User{}
	user3.Id = 3
	user3.FirstName = "Matheus"
	user3.LastName = "Pacheco"
	user3.Email = "mp@gmail.com"
	if err := user3.SetAge(-1); err != nil {
		fmt.Println("[main] houve um error na atribuição da idade:", err.Error())
	}

	fmt.Printf("User 1: %+v. É adulto?: %v\n", user1, user1.IsAdult())
	fmt.Printf("User 2: %+v. É adulto?: %v\n", user2, user2.IsAdult())

	objJSON, _ := json.Marshal(user1)

	fmt.Printf("JSON format: %v\n", string(objJSON))
}