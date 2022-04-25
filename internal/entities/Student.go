package entities

import "fmt"

type Student struct {
	//the id
	Id int `json:"Id" validate:"required,min=1,max=8,alpha_space"`
	//the student's first name
	FirstName string `json:"FirstName" validate:"required,min=1,max=30,alpha_space"`
	//the student's last name
	LastName string `json:"LastName" validate:"required,min=1,max=30,alpha_space"`
	//the student's age
	Age int `json:"Age" validate:"required,min=1,max=3,alpha_space"`
	//the codes of the languages
	LanguageCode string `json:"LanguageCode" validate:"required,min=2,max=50,alpha_space"`
}

func NewStudent(id int, firstname string, name string, age int, code string) Student {
	return Student{
		Id:           id,
		FirstName:    firstname,
		LastName:     name,
		Age:          age,
		LanguageCode: code,
	}
}

func (s Student) String() string {
	return fmt.Sprintf("Id: %d \nFirstName: %s \nLastName: %s \nAge: %d "+
		"\nCode: %s", s.Id, s.FirstName, s.LastName, s.Age, s.LanguageCode)
}
