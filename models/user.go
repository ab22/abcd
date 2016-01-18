package models

import (
	"time"
)

type User struct {
	Id        int
	Username  string `sql:"size:30; unique_index; not null"`
	Password  string
	Email     string `sql:"size:60; unique_index"`
	FirstName string `sql:"size:60"`
	LastName  string `sql:"size:60"`
	Status    int
	IsAdmin   bool
	IsTeacher bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
