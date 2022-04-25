package persistance

import . "Tp_Rest/internal/entities"

type StudentsDAO interface {
	FindAll() map[int]Student
	Find(id int) Student
	Exists(id int) bool
	Delete(id int) bool
	Create(student Student) bool
	Update(student Student) bool
}
