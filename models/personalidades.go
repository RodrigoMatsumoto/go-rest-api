package models

type Personalidade struct {
	Id       int    `jason:"id"`
	Nome     string `jason:"nome"`
	Historia string `jason:"historia"`
}

var Personalidades []Personalidade
