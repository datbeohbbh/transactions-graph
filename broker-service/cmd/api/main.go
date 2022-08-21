package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	webPort = os.Getenv("WEB_PORT")
)

type Broker struct {
	grpcConn *grpc.ClientConn
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

	broker := Broker{}
	conn, err := connectToGrpcServer(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	broker.grpcConn = conn

	r.POST("/broker", broker.routes)

	r.Run(fmt.Sprintf(":%s", webPort))
}

func connectToGrpcServer(ctx context.Context) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	SERVER_NAME := os.Getenv("GRPC_SERVER")
	NGINX_GRPC_PORT := os.Getenv("NGINX_GRPC_PORT")
	log.Printf("connect to nginx grpc load balancer %s", NGINX_GRPC_PORT)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%s", SERVER_NAME, NGINX_GRPC_PORT), opts...)
	if err != nil {
		return nil, fmt.Errorf("failed on connect to grpc server: %v", err)
	}
	return conn, nil
}
