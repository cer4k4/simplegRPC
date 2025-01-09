package main

import (
	"context"
	"log"
	"time"

	pb "grpc-example/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBybitClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	deposites, err := c.GetListOfDeposite(ctx, &pb.GetListOfDepositeRequest{})

	log.Println("response", deposites)
}
