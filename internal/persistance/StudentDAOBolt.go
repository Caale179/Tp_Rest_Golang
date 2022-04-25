package persistance

import (
	. "Tp_Rest/internal/entities"
	"encoding/json"
	"log"
	"sort"
	"strconv"
)

type StudentDAOBolt struct {
	DbBolt BoltDb
}

func NewStudentDAOBolt() StudentDAOBolt {
	return StudentDAOBolt{}
}

func (s StudentDAOBolt) Find(id int) Student {
	studentString := s.DbBolt.dbGet("students", strconv.Itoa(id))
	student := new(Student)
	decodeErr := json.Unmarshal([]byte(studentString), &student)
	if decodeErr != nil {
		log.Fatal("Find Unmarshalling error" + decodeErr.Error())
	}
	return *student
}

func (s StudentDAOBolt) FindAll() map[int]Student {
	studentsStrings := s.DbBolt.dbGetAll("students")
	students := make(map[int]Student)

	for _, element := range studentsStrings {
		student := new(Student)
		decodeErr := json.Unmarshal([]byte(element), &student)
		if decodeErr != nil {
			log.Fatal("FindAll Unmarshalling error" + decodeErr.Error())
		}
		students[student.Id] = *student
	}

	keys := make([]int, 0, len(students))
	for k := range students {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return students
}

func (s *StudentDAOBolt) Exists(id int) bool {
	return s.Find(id) != Student{}
}

func (s *StudentDAOBolt) Delete(id int) bool {
	exists := s.Exists(id)
	if exists {
		s.DbBolt.dbDelete("students", strconv.Itoa(id))
	}
	return exists
}

func (s *StudentDAOBolt) Create(student Student) bool {
	exists := s.Exists(student.Id)
	if !exists {
		encoded, encodedErr := json.Marshal(student)
		if encodedErr != nil {
			log.Fatal("Create Marshalling error" + encodedErr.Error())
		}
		s.DbBolt.DbPut("students", strconv.Itoa(student.Id), string(encoded))
	}

	return !exists
}

func (s *StudentDAOBolt) Update(student Student) bool {
	exists := s.Exists(student.Id)
	if exists {
		encoded, encodedErr := json.Marshal(student)
		if encodedErr != nil {
			log.Fatal("Update Marshalling error" + encodedErr.Error())
		}
		s.DbBolt.DbPut("students", strconv.Itoa(student.Id), string(encoded))
	}

	return exists
}
