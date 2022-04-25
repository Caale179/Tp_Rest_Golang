package entities

import "fmt"

//the programing language
type Language struct {
	// Code of the language
	// in: string
	Code string `json:"Code" validate:"required,min=2,max=2,alpha_space"`
	// Name of the language
	// in: string
	Name string `json:"Name" validate:"required,min=2,max=20,alpha_space"`
}

func NewLanguage(code string, name string) Language {
	return Language{
		Code: code,
		Name: name,
	}
}

func (l Language) String() string {
	return fmt.Sprintf("Code: %s \nName: %s", l.Code, l.Name)
}
