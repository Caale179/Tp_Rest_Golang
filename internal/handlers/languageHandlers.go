package handlers

import (
	"Tp_Rest/internal/entities"
	"Tp_Rest/internal/persistance"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var daoLanguage = persistance.NewLanguageDAOBolt()

type HandlerL struct {
	LanguageDao persistance.LanguagesDAO
}

// swagger:operation GET /languages/{code} Language LanguageHandler
// ---
// summary: Returns the language corresponding to the code
// parameters:
// - name: id
//   in: path
//   description: language code
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/languageRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (l *HandlerL) LanguageHandler(w http.ResponseWriter, r *http.Request) {
	//url := r.URL.String()
	//index := url[len(url)-1:]
	//leLanguage := daoLanguage.Find(index)
	//
	//fmt.Fprintf(w, "<html><head></head><body>")
	//fmt.Fprintf(w, "<p>Code: %s   | Name: %s</p> <br>", leLanguage.Code, leLanguage.Name)
	//fmt.Fprintf(w, "</body></html>")
	vars := mux.Vars(r)
	code := vars["code"]

	jsonString, err := json.Marshal(l.LanguageDao.Find(code))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(jsonString))
}

// swagger:operation GET /languages Language AllLanguageHandler
// ---
// summary: Returns all the known languages
// responses:
//   "200":
//     "$ref": "#/responses/languagesRes"
//   "400":
//     "$ref": "#/responses/badReq"
func (l *HandlerL) AllLanguageHandler(w http.ResponseWriter, _ *http.Request) {
	//	fmt.Fprintf(w, "<html><head></head><body>")
	//	for _, lang := range daoLanguage.FindAll() {
	//		fmt.Fprintf(w, "<p>Code: %s   | Name: %s</p> <br>", lang.Code, lang.Name)
	//	}
	//	fmt.Fprintf(w, "</body></html>")
	jsonString, err := json.Marshal(l.LanguageDao.FindAll())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(jsonString))
}

// swagger:operation POST /languages Language CreateLanguageHandler
// ---
// summary: Adds the new given language to the database
// parameters:
// - name: language
//   description: the language to add
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Language"
// responses:
//   "201":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
func (l *HandlerL) CreateLanguageHandler(w http.ResponseWriter, r *http.Request) {
	//reqBody, _ := ioutil.ReadAll(r.Body)
	//var language entities.Language
	//err := json.Unmarshal(reqBody, &language)
	//if err != nil {
	//	return
	//}
	//err = json.NewEncoder(w).Encode(language)
	//if err != nil {
	//	return
	//}
	//newLanguage := daoLanguage.Create(language)
	//fmt.Fprintf(w, "new lang : %s", newLanguage)
	//
	//newData, err := json.Marshal(language)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(string(newData))
	//}

	language := entities.NewLanguage("", "")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &language)

	isOk, err := json.Marshal(l.LanguageDao.Create(language))

	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(201)
	fmt.Fprintf(w, string(isOk))
}

// swagger:operation PUT /languages Language UpdateLanguageHandler
// ---
// summary: Updates the caracteristics of a programming language
// parameters:
// - name: language
//   description: the language to update
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Language"
// responses:
//   "200":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (l *HandlerL) UpdateLanguageHandler(w http.ResponseWriter, r *http.Request) {
	//reqBody, _ := ioutil.ReadAll(r.Body)
	//var language entities.Language
	//err := json.Unmarshal(reqBody, &language)
	//if err != nil {
	//	return
	//}
	//err = json.NewEncoder(w).Encode(language)
	//if err != nil {
	//	return
	//}
	//newLanguage := daoLanguage.Update(language)
	//fmt.Fprintf(w, "new lang : %s", newLanguage)
	//
	//newData, err := json.Marshal(language)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(string(newData))
	//}
	language := entities.NewLanguage("", "")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &language)

	isOk, err := json.Marshal(l.LanguageDao.Update(language))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(isOk))
}

// swagger:operation DELETE /languages/{code} Language DeleteLanguageHandler
// ---
// summary: Supprime le langage ayant le code spécifié
// parameters:
// - name: id
//   in: path
//   description: the code of the language to delete
//   type: string
//   required: true
// - name: language
//   description: The language to delete
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Language"
// responses:
//   "200":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (l *HandlerL) DeleteLanguageHandler(w http.ResponseWriter, r *http.Request) {
	//url := r.URL.String()
	//index := url[len(url)-1:]
	//if daoLanguage.Delete(index) {
	//	fmt.Fprintf(w, "good")
	//	return
	//}
	//fmt.Fprintf(w, "nop")
	vars := mux.Vars(r)
	code := vars["code"]

	isOk, err := json.Marshal(l.LanguageDao.Delete(code))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(isOk))
}
