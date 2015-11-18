package services

import (
	"fmt"

	"github.com/ab22/abcd/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	db gorm.DB

	// Define and export global service instances to avoid
	// creating new instances everytime it is needed.
	AuthService = authService{}
	UserService = userService{}
)

// Creates a connection to the database and assigns it
// to the global DB instance so that every service
// can access it.
func Initialize() error {
	var err error
	dbCfg := config.EnvVariables.Db

	connectionString := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Name,
	)

	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	db.LogMode(config.DbLogMode)

	return err
}
