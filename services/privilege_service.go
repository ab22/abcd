package services

import (
	"github.com/ab22/abcd/models"
)

// Contains all of the logic for user privileges.
type privilegeService struct{}

// Returns all privileges for a Role.
func (s *privilegeService) GetAllByRoleId(roleId int) ([]models.Privilege, error) {
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
		Find(&privileges)

	if err != nil {
		if err != gorm.RecordNotFound {
			return nil, err
		}
	}

	return privileges, nil
}
