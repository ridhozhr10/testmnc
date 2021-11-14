package models

import "errors"

var globalLanguange = []Languange{}

type Languange struct {
	Languange  string   `json:"languange"`
	Appeared   int      `json:"appeared"`
	Created    []string `json:"created"`
	Functional bool     `json:"functional"`
	Relation   Relation `json:"relation"`
}

type Relation struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences   []string `json:"influences"`
}

func RemoveLanguage(index int) error {
	if len(globalLanguange) > index {
		return errors.New("invalid index")
	}
	globalLanguange = append(globalLanguange[:index], globalLanguange[index+1:]...)
	return nil
}

func UpdateLanguage(index int, lang Languange) {
	globalLanguange[index] = lang
}

func AddLanguage(newLang Languange) {
	globalLanguange = append(globalLanguange, newLang)
}

func GetLanguage(index int) Languange {
	return globalLanguange[index]
}

func GetAllLanguage() []Languange {
	return globalLanguange
}

func InitLanguage() {
	newLang := Languange{
		Languange:  "C",
		Appeared:   1972,
		Created:    []string{"Dennis Ritchie"},
		Functional: true,
		Relation: Relation{
			InfluencedBy: []string{
				"B",
				"ALGOL 68",
				"Assembly",
				"FORTRAN",
			},
			Influences: []string{
				"C++",
				"Objective-C",
				"C#",
				"Java",
				"Javascript",
				"PHP",
				"Go",
			},
		},
	}

	globalLanguange = append(globalLanguange, newLang)
}
