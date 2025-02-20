package person

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	Person
}

func PromoteToAdmin(email string, p Person) Admin {
	return Admin{
		email:  email,
		Person: p,
	}
}

// reciever argument
func (p Person) Show() {
	fmt.Println(p.firstName)
}

func (p *Person) CapitalizeFirstName() {
	// just for struct you can direct access the value instead (*u).firstName
	p.firstName = strings.ToUpper(p.firstName)

}

func New(firstName, lastName, birthdate string) (*Person, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("please fill all the fields")
	}

	return &Person{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
