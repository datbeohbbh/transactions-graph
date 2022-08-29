package main

import (
	"context"
	"fmt"
	"os"

	addressManager "github.com/datbeohbbh/transactions-graph/broker/address-manager"

	graph "github.com/datbeohbbh/transactions-graph/broker/graph-client"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	webPort = os.Getenv("WEB_PORT")
)

type Broker struct {
	addressManagerClient *addressManager.Client

	graphClient *graph.GraphClient
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	conn, err := connectToGRPCServerContext(context.Background(),
		os.Getenv("NGINX_SERVICE_NAME"),
		os.Getenv("NGINX_GRPC_PORT"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())

	if err != nil {
		panic(fmt.Errorf("failed on connect to nginx: %v", err))
	}
	defer conn.Close()

	broker := Broker{
		addressManagerClient: addressManager.New(conn),
		graphClient:          graph.New(conn),
	}

	r.POST("/broker", broker.routes)

	r.Run(fmt.Sprintf(":%s", webPort))
}
