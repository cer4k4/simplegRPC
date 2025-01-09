package main

import (
	"context"
	"log"
	"net"

	httpbybit "grpc-example/http"
	pb "grpc-example/proto"

	"google.golang.org/grpc"
)

type bybit struct {
	HttpBybit *httpbybit.BybitHttp
	pb.UnimplementedBybitServer
}

func (s *bybit) GetAllAssetWithNetworks(ctx context.Context, req *pb.GetAllAssetWithNetworksRequest) (*pb.GetAllAssetWithNetworksRespose, error) {
	if err, result := s.HttpBybit.GetCoinInfo(); err != nil {
		log.Println(err)
	} else {

		return &pb.GetAllAssetWithNetworksRespose{append(result.Rows, result.Rows...)}
	}

}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	apiSecret := "NyPYv6iE43RIJCuMMVVbBBgQh7PGvW8Kd1gR"
	apiKey := "2B4Lk5MQuA8P5CLc78"

	httpBybit := httpbybit.NewBybitHttp(apiKey, apiSecret)

	s := grpc.NewServer()

	pb.RegisterBybitServer(s, &bybit{HttpBybit: httpBybit})
	log.Println("Server starting on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
