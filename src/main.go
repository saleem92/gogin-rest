package main

import (
	"github.com/gin-gonic/gin"

	"go-rest/basics/controllers"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Print(".env file not found")
	// }

	// ConnectDB()

	gin.DisableConsoleColor()

	router := gin.Default()

	controllers.SetupRoutes(router)
	router.Run()
}
