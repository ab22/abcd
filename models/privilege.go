package models

type Privilege struct {
	Key         string `gorm:"primary_key"; sql:"size:20"; unique_index`
	Description string `sql:"size:200"`
	RoleId      int

	Role Role
}
