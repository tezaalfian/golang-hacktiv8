package database

import (
	"fmt"
	"go-auth-product/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "root"
	dbname   = "go_gin_product"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Database connected successfully")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
