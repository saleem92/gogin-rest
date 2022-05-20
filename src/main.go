package main

import (
	"github.com/gin-gonic/gin"

	"go-rest/basics/controllers"
	"go-rest/basics/messaging"

	dapr "github.com/dapr/go-sdk/client"
)

func main() {

	gin.DisableConsoleColor()

	router := gin.Default()
	go messaging.CreateSubscriber()
	go createDaprClient(router)

	controllers.SetupRoutes(router)
	router.Run()
}

func createDaprClient(r *gin.Engine) {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}

	messaging.SetupPublishRoute(r, client)
}
