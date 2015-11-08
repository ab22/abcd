package models

type Role struct {
	Id   int
	Name string `sql:"size:30"`

	Privileges []Privilege
}
