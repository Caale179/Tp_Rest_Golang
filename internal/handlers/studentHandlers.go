package handlers

import (
	. "Tp_Rest/internal/entities"
	"Tp_Rest/internal/persistance"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

var daoStudent = persistance.NewStudentDAOBolt()

type HandlerS struct {
	StudentDao persistance.StudentsDAO
}

// swagger:operation GET /students/{id} Student StudentHandler
// ---
// summary: Returns the given student via his Id
// parameters:
// - name: id
//   in: path
//   description: the student's Id
//   type: int
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/studentRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (s *HandlerS) StudentHandler(w http.ResponseWriter, r *http.Request) {
	//url := r.URL.String()
	//id := url[len(url)-1:]
	//idInt, _ := strconv.Atoi(id)
	//leStudent := daoStudent.Find(idInt)
	//
	//fmt.Fprintf(w, "<html><head></head><body>")
	//fmt.Fprintf(w, "<p>Id: %d | FullName: %s %s | Age: %d</p> <br>", leStudent.Id, leStudent.LastName, leStudent.FirstName, leStudent.Age)
	//fmt.Fprintf(w, "</body></html>")
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonString, err := json.Marshal(s.StudentDao.Find(idInt))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(jsonString))
}

// swagger:operation GET /students Student AllStudentHandler
// ---
// summary: returns the list of all the known students
// responses:
//   "200":
//     "$ref": "#/responses/studentsRes"
//   "400":
//     "$ref": "#/responses/badReq"
func (s *HandlerS) AllStudentHandler(w http.ResponseWriter, _ *http.Request) {
	//fmt.Fprintf(w, "<html><head></head><body>")
	//for _, leStudent := range daoStudent.FindAll() {
	//	fmt.Fprintf(w, "<p>Id: %d | FullName: %s %s | Age: %d</p> <br>", leStudent.Id, leStudent.LastName, leStudent.FirstName, leStudent.Age)
	//}
	//fmt.Fprintf(w, "</body></html>")
	jsonString, err := json.Marshal(s.StudentDao.FindAll())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(jsonString))
}

// swagger:operation POST /students Student CreateStudentHandler
// ---
// summary: Adds a new student
// parameters:
// - name: student
//   description: The new student to add
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   "201":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
func (s *HandlerS) CreateStudentHandler(w http.ResponseWriter, r *http.Request) {

	//reqBody, _ := ioutil.ReadAll(r.Body)
	//var student Student
	//err := json.Unmarshal(reqBody, &student)
	//if err != nil {
	//	return
	//}
	//err = json.NewEncoder(w).Encode(student)
	//if err != nil {
	//	return
	//}
	//newstudent := daoStudent.Create(student)
	//fmt.Fprintf(w, "new lang : %s", newstudent)
	//
	//newData, err := json.Marshal(student)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(string(newData))
	//}
	student := NewStudent(0, "", "", 0, "")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &student)

	isOk, err := json.Marshal(s.StudentDao.Create(student))

	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(201)
	fmt.Fprintf(w, string(isOk))
}

// swagger:operation PUT /students Student UpdateStudentHandler
// ---
// summary: Updates the given student
// parameters:
// - name: student
//   description: the student to update
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   "200":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (s *HandlerS) UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	//reqBody, _ := ioutil.ReadAll(r.Body)
	//var student Student
	//err := json.Unmarshal(reqBody, &student)
	//if err != nil {
	//	return
	//}
	//err = json.NewEncoder(w).Encode(student)
	//if err != nil {
	//	return
	//}
	//newstudent := daoStudent.Update(student)
	//fmt.Fprintf(w, "new lang : %s", newstudent)
	//
	//newData, err := json.Marshal(student)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(string(newData))
	//}
	student := NewStudent(0, "", "", 0, "")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &student)

	isOk, err := json.Marshal(s.StudentDao.Update(student))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(isOk))
}

// swagger:operation DELETE /students/{id} Student DeleteStudentHandler
// ---
// summary: Deletes the given student
// parameters:
// - name: id
//   in: path
//   description: the student to delete's Id
//   type: string
//   required: true
// - name: student
//   description: The student to delete
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   "200":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (s *HandlerS) DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	//url := r.URL.String()
	//id := url[len(url)-1:]
	//idInt, _ := strconv.Atoi(id)
	//if !daoStudent.Delete(idInt) {
	//	fmt.Fprintf(w, "nop")
	//}
	//fmt.Fprintf(w, "good")
	//return
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	isOk, err := json.Marshal(s.StudentDao.Delete(idInt))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(isOk))
}
