package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-rest/basics/controllers"
	"go-rest/basics/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
)

func main() {
	ConnectDB()

	gin.DisableConsoleColor()

	// Used to add file logger
	// file, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(file)

	router := gin.Default()

	controllers.SetupRoutes(router)

	router.Run()
}

func ConnectDB() {
	dbName := "test_db"

	createIfDBNotExists(dbName)
	connectionString := fmt.Sprintf("postgres://postgres:root@localhost:5432/%s", dbName)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("failed to connect to DB")
	}

	db.AutoMigrate(&models.Album{})
}

func createIfDBNotExists(dbName string) {
	defaultDB, err := gorm.Open(postgres.Open("postgres://postgres:root@localhost:5432/postgres"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to DB")
	}

	var result string
	defaultDB.Raw("SELECT datname FROM pg_database where datname=" + dbName).Scan(&result)

	if len(result) == 0 {
		defaultDB.Exec("CREATE DATABASE " + dbName)

		if defaultDB.Error != nil {
			panic("failed to create %s database")
		}
	}
}
