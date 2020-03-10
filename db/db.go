package db

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB() (*gorm.DB, error) {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())
	if err != nil {
		fmt.Println("Database error: ", err)
	}
	db.DB().SetMaxIdleConns(2)
	db.LogMode(true)
	return db, err
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		// Enter models here, like this:
		// &model.Ksher{},
	)
}

