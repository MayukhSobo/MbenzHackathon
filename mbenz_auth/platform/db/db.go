package db

import (
	"github.com/jinzhu/gorm"
	"os"
)

type Database struct {
	DB *gorm.DB
}


//func (db *Database) IfReachable() {
//
//}


//func (db *Database) Close() error {
//	err := db.DB.Close()
//	if err != nil {
//		return nil
//	}
//	return fmt.Errorf("unable to close the database")
//}

// NewDatabase creates a new database of a type
func NewDatabase() (*Database, error){
	return NewSQLConn(os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))
}
