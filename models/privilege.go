package models

type Privilege struct {
	Id          int
	Key         string `sql:"size:20"; unique_index`
	Description string `sql:"size:200"`
	RoleId      int

	Role Role
}
