package messaging

import (
	"context"
	"fmt"
	"io"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/gin-gonic/gin"
)

func SetupPublishRoute(r *gin.Engine, client dapr.Client) {
	r.POST("/publish", handlePublish(client))
}

func handlePublish(client dapr.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		data, _ := io.ReadAll(c.Request.Body)

		if err := client.PublishEvent(ctx, "pubsub", "orders", data); err != nil {
			fmt.Println("Error occured while publishing", err)
			return
		}

		fmt.Println("Published data: " + string(data))

		defer client.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Message published "})
	}
}
