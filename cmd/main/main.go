package main

import (
	"Tp_Rest/internal/entities"
	. "Tp_Rest/internal/handlers"
	"Tp_Rest/internal/persistance"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

//  Tim Beziau:
//   version: 0.0.1
//   title: The TimPI
//  Schemes: http, https
//  Host: localhost:8081
//  BasePath: /
//  Produces:
//    - application/json
//
// securityDefinitions:
// swagger:meta
func main() {
	//ouverture de la bdd
	var l HandlerL

	var s HandlerS

	r := mux.NewRouter()

	bdd := new(persistance.BoltDb)
	bdd.DbOpen("bdd.db")
	defer bdd.DbClose()
	bdd.DbCreateBucket("languages")
	bdd.DbCreateBucket("students")

	//Students
	leStudentDAO := new(persistance.StudentDAOBolt)

	student1 := entities.NewStudent(1, "Michel", "Baie", 25, "js")
	student2 := entities.NewStudent(2, "Paul", "Patine", 19, "c")

	encoded, encodedErr := json.Marshal(student1)
	if encodedErr != nil {
		log.Fatal(encodedErr)
	}
	bdd.DbPut("students", strconv.Itoa(student1.Id), string(encoded))

	encoded, encodedErr = json.Marshal(student2)
	if encodedErr != nil {
		log.Fatal(encodedErr)
	}
	bdd.DbPut("students", strconv.Itoa(student2.Id), string(encoded))

	leStudentDAO.DbBolt = *bdd

	//Language
	//Languages
	leLanguageDAO := new(persistance.LanguageDAOBolt)

	language1 := entities.NewLanguage("js", "JavaScript")
	language2 := entities.NewLanguage("c", "C")

	encoded, encodedErr = json.Marshal(language1)
	if encodedErr != nil {
		log.Fatal(encodedErr)
	}
	bdd.DbPut("languages", language1.Code, string(encoded))

	encoded, encodedErr = json.Marshal(language2)
	if encodedErr != nil {
		log.Fatal(encodedErr)
	}
	bdd.DbPut("languages", language2.Code, string(encoded))

	leLanguageDAO.DbBolt = *bdd
	l.LanguageDao = leLanguageDAO
	s.StudentDao = leStudentDAO
	r.HandleFunc("/rest/languages/{code}", l.LanguageHandler).Methods("GET")
	r.HandleFunc("/rest/languages", l.AllLanguageHandler).Methods("GET")
	r.HandleFunc("/rest/languages", l.CreateLanguageHandler).Methods("POST")
	r.HandleFunc("/rest/languages", l.UpdateLanguageHandler).Methods("PUT")
	r.HandleFunc("/rest/languages", l.DeleteLanguageHandler).Methods("DELETE")

	r.HandleFunc("/rest/students/{id}", s.StudentHandler).Methods("GET")
	r.HandleFunc("/rest/students", s.AllStudentHandler).Methods("GET")
	r.HandleFunc("/rest/students", s.CreateStudentHandler).Methods("POST")
	r.HandleFunc("/rest/students", s.UpdateStudentHandler).Methods("PUT")
	r.HandleFunc("/rest/students", s.DeleteStudentHandler).Methods("DELETE")

	//swagger documentation
	fs := http.FileServer(http.Dir("./swagger"))
	r.PathPrefix("/doc/").Handler(http.StripPrefix("/doc/", fs))

	err := http.ListenAndServe(":8081", r)

	if err != nil {
		log.Fatal(err)
	}
}
