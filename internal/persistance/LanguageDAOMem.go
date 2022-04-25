package persistance

//
//import (
//	. "Tp_Rest/internal/entities"
//	"sort"
//)
//
//type LanguageDAOMem struct{}
//
//var mapLanguage = make(map[string]Language)
//var _ LanguagesDAO = (*LanguageDAOMem)(nil)
//
//func NewLanguageDAOMemory() LanguageDAOMem {
//	return LanguageDAOMem{}
//}
//
//func (l LanguageDAOMem) FindAll() []Language {
//	var languages []Language
//	for _, value := range mapLanguage {
//		languages = append(languages, value)
//	}
//	sort.Slice(languages, func(i, j int) bool {
//		return languages[i].Code < languages[j].Code
//	})
//	return languages
//}
//
//func (l LanguageDAOMem) Find(id string) *Language {
//	language, exists := mapLanguage[id]
//	if exists {
//		return &language
//	} else {
//		return nil
//	}
//}
//
//func (l LanguageDAOMem) Exists(id string) bool {
//	_, exists := mapLanguage[id]
//	if exists {
//		return true
//	} else {
//		return false
//	}
//}
//
//func (l LanguageDAOMem) Delete(id string) bool {
//	if l.Exists(id) {
//		return false
//	} else {
//		delete(mapLanguage, id)
//		return true
//	}
//}
//
//func (l LanguageDAOMem) Create(language Language) bool {
//	if l.Exists(language.Code) {
//		return false
//	} else {
//		mapLanguage[language.Code] = language
//		return true
//	}
//}
//
//func (l LanguageDAOMem) Update(language Language) bool {
//	if l.Exists(language.Code) {
//		mapLanguage[language.Code] = language
//		return true
//	} else {
//		return false
//	}
//}
