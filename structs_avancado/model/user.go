package model

// User model that represents user attributes
type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	age       int    `json:"age"`
	isAdult   bool   `json:"is_adult"`
}

func (u *User) SetAge(age int) {
	u.age = age

	if age >= 18 {
		u.isAdult = true
	} else {
		u.isAdult = false
	}
}

func (u *User) GetAge() int { return u.age }

func (u *User) IsAdult() bool { return u.isAdult }
