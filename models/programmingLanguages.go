package models

type ProgrammingLangs struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Typed   bool   `json:"typed"`
	History string `json:"history"`
}

var ProgrammingLanguages []ProgrammingLangs
