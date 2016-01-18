package models

import (
	"time"
)

type Student struct {
	Id           int    // Database autogenerated Id number.
	IdNumber     string `sql:"not null; size:40; unique_index"` // Honduran Document ID Number or Passport Number.
	Email        string `sql:"size:60; unique_index"`
	FirstName    string `sql:"size:60"`
	LastName     string `sql:"size:60"`
	Status       int
	PlaceOfBirth string
	Address      string
	Birthdate    time.Time
	Gender       rune
	Nationality  string
	PhoneNumber  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
