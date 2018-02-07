package model

import "errors"

// User model that represents user attributes
type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	age       int    `json:"age"`
	isAdult   bool   `json:"is_adult"`
}

var ErrInvalidAge = errors.New("Age cannot be negative")

func (u *User) SetAge(age int) (err error) {
	err = nil

	if age < 0 {
		err = ErrInvalidAge
		return
	}

	u.age = age

	if age >= 18 {
		u.isAdult = true
	} else {
		u.isAdult = false
	}

	return
}

func (u *User) GetAge() int { return u.age }

func (u *User) IsAdult() bool { return u.isAdult }
