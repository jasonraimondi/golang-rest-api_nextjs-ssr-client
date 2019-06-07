package models

type Person struct {
	FirstName *string `db:"first_name"`
	LastName  *string `db:"last_name"`
	Email     string
}

func NewSimplePerson(email string) (p Person) {
	return Person{
		Email:     email,
		FirstName: nil,
		LastName:  nil,
	}
}