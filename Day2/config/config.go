package config

import (
	"Day2/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	config := map[string]string{
		"DB_username": "debian-sys-maint",
		"DB_password": "7UfuSJnEvNkWOW7D",
		"DB_port":     "3306",
		"DB_host":     "127.0.0.1",
		"DB_name":     "test_db",
	}

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config["DB_username"],
			config["DB_password"],
			config["DB_host"],
			config["DB_port"],
			config["DB_name"])

	// fmt.Println(connectionString)
	var e error
	database, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	fmt.Print(DB)
	if e != nil {
		// panic(e)
		os.Exit(0)
	}

	// database.AutoMigrate(&models.User{})

	DB = database

	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(&models.User{})
}
