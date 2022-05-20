package messaging

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

func CreateSubscriber() {
	var sub = &common.Subscription{
		PubsubName: "pubsub",
		Topic:      "orders",
		Route:      "/checkout",
	}

	s := daprd.NewService(":6002")
	//Subscribe to a topic
	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		fmt.Printf("error adding topic subscription: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("error listenning: %v", err)
	}
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	fmt.Println("Subscriber received: ", e.Data)
	return false, nil
}
