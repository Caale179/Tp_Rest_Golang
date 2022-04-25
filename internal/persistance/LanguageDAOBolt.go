package persistance

import (
	. "Tp_Rest/internal/entities"
	"encoding/json"
	"log"
	"sort"
)

type LanguageDAOBolt struct {
	DbBolt BoltDb
}

func NewLanguageDAOBolt() LanguageDAOBolt {
	return LanguageDAOBolt{}
}

func (l LanguageDAOBolt) FindAll() map[string]Language {
	languagesStrings := l.DbBolt.dbGetAll("languages")
	languages := make(map[string]Language)

	for _, element := range languagesStrings {
		language := new(Language)
		decodeErr := json.Unmarshal([]byte(element), &language)
		if decodeErr != nil {
			log.Fatal(decodeErr)
		}
		languages[language.Code] = *language
	}

	keys := make([]string, 0, len(languages))
	for k := range languages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return languages
}

func (l *LanguageDAOBolt) Find(code string) Language {
	languageString := l.DbBolt.dbGet("languages", code)
	language := new(Language)
	decodeErr := json.Unmarshal([]byte(languageString), &language)
	if decodeErr != nil {
		log.Fatal("Find Unmarshalling error" + decodeErr.Error())
	}
	return *language
}

func (l *LanguageDAOBolt) Exists(code string) bool {
	exists := l.Find(code) != Language{}

	return exists
}

func (l *LanguageDAOBolt) Delete(code string) bool {
	exists := l.Exists(code)
	if exists {
		l.DbBolt.dbDelete("languages", code)
	}

	return exists
}

func (l *LanguageDAOBolt) Create(language Language) bool {
	exists := l.Exists(language.Code)
	if !exists {
		encoded, encodedErr := json.Marshal(language)
		if encodedErr != nil {
			log.Fatal("Create Marshalling error" + encodedErr.Error())
		}
		l.DbBolt.DbPut("languages", language.Code, string(encoded))
	}

	return !exists
}

func (l *LanguageDAOBolt) Update(language Language) bool {
	exists := l.Exists(language.Code)
	if exists {
		encoded, encodedErr := json.Marshal(language)
		if encodedErr != nil {
			log.Fatal("Update Marshalling error" + encodedErr.Error())
		}
		l.DbBolt.DbPut("languages", language.Code, string(encoded))
	}

	return exists
}
