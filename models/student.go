package models

import (
	"time"
)

type Student struct {
	Id        int
	Email     string `sql:"size:60; unique_index"`
	FirstName string `sql:"size:60"`
	LastName  string `sql:"size:60"`
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
