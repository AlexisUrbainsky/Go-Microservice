package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectToDB() (*gorm.DB, error) {

	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	dns := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s sslmode=disable", host, name, user, password, port)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: CreateLogsConfigGorm(),
	})

	if err != nil {
		log.Printf("failed to connect database")
	}

	return db, err

}
