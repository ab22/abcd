package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github/ab22/abcd/config"
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

	return nil
}
