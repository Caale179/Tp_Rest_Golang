package persistance

import . "Tp_Rest/internal/entities"

type LanguagesDAO interface {
	FindAll() map[string]Language
	Find(id string) Language
	Exists(id string) bool
	Delete(id string) bool
	Create(language Language) bool
	Update(language Language) bool
}
