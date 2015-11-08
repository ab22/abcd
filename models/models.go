package models

import (
	"fmt"

	"github.com/ab22/abcd/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Migrate() error {
	dbCfg := config.EnvVariables.Db
	connectionString := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Name,
	)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&Privilege{})

	return nil
}
