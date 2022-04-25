package persistance

//
//import (
//	. "Tp_Rest/internal/entities"
//	"fmt"
//	"sort"
//)
//
//type StudentDAOMem struct{}
//
//var mapStudents = make(map[int]Student)
//var _ StudentsDAO = (*StudentDAOMem)(nil)
//
//func NewStudentDAOMemory() StudentDAOMem {
//	return StudentDAOMem{}
//}
//
//func (s StudentDAOMem) Find(id int) *Student {
//	st, exists := mapStudents[id]
//	if exists {
//		return &st
//	} else {
//		return nil
//	}
//}
//
//func (s StudentDAOMem) FindAll() []Student {
//	var students []Student
//	for _, value := range mapStudents {
//		students = append(students, value)
//	}
//	sort.Slice(students, func(i, j int) bool {
//		return students[i].Id < students[j].Id
//	})
//	return students
//}
//
//func (s StudentDAOMem) Exists(id int) bool {
//	_, exists := mapStudents[id]
//	if exists {
//		return true
//	} else {
//		return false
//	}
//}
//
//func (s StudentDAOMem) Delete(id int) bool {
//	if s.Exists(id) {
//		return false
//	} else {
//		delete(mapStudents, id)
//		return true
//	}
//}
//
//func (s StudentDAOMem) Create(student Student) bool {
//	fmt.Printf("ca marche Create")
//	if s.Exists(student.Id) || student.Id <= 0 {
//		return false
//	} else {
//		mapStudents[student.Id] = student
//		return true
//	}
//}
//
//func (s StudentDAOMem) Update(student Student) bool {
//	if s.Exists(student.Id) {
//		mapStudents[student.Id] = student
//		return true
//	} else {
//		return false
//	}
//}
