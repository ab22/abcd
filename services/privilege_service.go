package services

import (
	"github.com/ab22/abcd/models"
	"github.com/jinzhu/gorm"
)

// Contains all of the logic for user privileges.
type privilegeService struct{}

// Returns all privileges for a Role.
func (s *privilegeService) FindAllByRoleId(roleId int) ([]models.Privilege, error) {
	var privileges []models.Privilege
	var err error

	role, err := RoleService.Find(roleId)

	if err != nil {
		return nil, err
	} else if role == nil {
		return privileges, nil
	}

	err = db.Model(&role).
		Association("Privileges").
		Find(&privileges).
		Error

	if err != nil {
		if err != gorm.RecordNotFound {
			return nil, err
		}
	}

	return privileges, nil
}
