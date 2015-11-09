package services

import (
	"github.com/ab22/abcd/models"
	"github.com/jinzhu/gorm"
)

// Contains all of the logic for user roles.
type roleService struct{}

// Find a role by role id.
func (s *roleService) Find(roleId int) (*models.Role, error) {
	role := &models.Role{}

	err := db.Where("id = ?", roleId).
		First(role).
		Error

	if err != nil {
		if err != gorm.RecordNotFound {
			return nil, err
		}

		return nil, nil
	}

	return role, nil
}

// Find all roles in database.
func (s *roleService) FindAll() ([]models.Role, error) {
	var roles []models.Role
	var err error

	err = db.Find(&roles).Error

	if err != nil {
		if err != gorm.RecordNotFound {
			return nil, err
		}
	}

	return roles, nil
}
