package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	gRPCPort = os.Getenv("GRPC_PORT")
	webPort  = os.Getenv("WEB_PORT")
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
	SERVER_NAME := os.Getenv("GRPC_SERVICE")
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%s", SERVER_NAME, gRPCPort), opts...)
	if err != nil {
		return nil, errors.New("failed on connect to grpc server")
	}
	return conn, nil
}
