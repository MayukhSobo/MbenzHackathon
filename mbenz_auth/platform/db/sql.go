package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"mbenz_auth/app/models"
)

// NewSQLConn creates a new database connection of given types with the credentials provided
func NewSQLConn(DBDriver, DBUser, DBPassword, DBPort, DBHost, DBName string) (*Database, error){
	var database *Database
	switch DBDriver {
		case "postgres":
			URI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
				DBHost, DBPort, DBUser, DBName, DBPassword)
			db, err := gorm.Open(DBDriver, URI)
			if err != nil {
				return nil, fmt.Errorf("cannot connect to %s database", DBDriver)
			}else{
				database = &Database{DB: db}

			}
		case "mysql":
			return nil, fmt.Errorf("%s database support is coming soon", DBDriver)

		default:
			return nil, fmt.Errorf("%s database is not supported", DBDriver)
	}
	database.DB.Debug().AutoMigrate(&models.User{})
	return database, nil

}