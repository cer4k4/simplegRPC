package main

import (
	"context"
	"log"
	"net"

	pb "grpc-example/proto"
	httpbybit "grpc-example/server/http"

	"google.golang.org/grpc"
)

type bybit struct {
	HttpBybit *httpbybit.BybitHttp
	pb.UnimplementedBybitServer
}

func (s *bybit) GetAllAssetWithNetworks(ctx context.Context, req *pb.GetAllAssetWithNetworksRequest) (*pb.GetAllAssetWithNetworksRespose, error) {
	err, result := s.HttpBybit.GetCoinInfo()
	if err != nil {
		log.Println(err)
	}
	// Create a slice of pointers to store the rows
	var List []*pb.GetAllAssetWithNetworksRespose_Rows

	for c := range result.Rows {
		// Create a new row and get its address
		ass := &pb.GetAllAssetWithNetworksRespose_Rows{
			Coin:         result.Rows[c].Coin,
			Name:         result.Rows[c].Name,
			RemainAmount: result.Rows[c].RemainAmount,
		}

		for i := range result.Rows[c].Chains {
			chain := &pb.GetAllAssetWithNetworksRespose_Chains{
				Chain:                 result.Rows[c].Chains[i].Chain,
				ChainDeposit:          result.Rows[c].Chains[i].ChainDeposit,
				ChainType:             result.Rows[c].Chains[i].ChainType,
				ChainWithdraw:         result.Rows[c].Chains[i].ChainWithdraw,
				Confirmation:          result.Rows[c].Chains[i].Confirmation,
				DepositMin:            result.Rows[c].Chains[i].DepositMin,
				MinAccuracy:           result.Rows[c].Chains[i].MinAccuracy,
				WithdrawFee:           result.Rows[c].Chains[i].WithdrawFee,
				WithdrawMin:           result.Rows[c].Chains[i].WithdrawMin,
				WithdrawPercentageFee: result.Rows[c].Chains[i].WithdrawPercentageFee,
			}
			ass.Chains = append(ass.Chains, chain)
		}
		List = append(List, ass)

	}
	return &pb.GetAllAssetWithNetworksRespose{Rows: List}, nil
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
