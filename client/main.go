package main

import (
	"context"
	"log"
	"time"

	"grpc-example/client/config"
	databases "grpc-example/client/database"
	pb "grpc-example/proto"

	"github.com/wuhewuhe/bybit.go.api/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	config := config.Loader()

	mongoDB := databases.NewMongoDB(config.MongoClient)
	c := pb.NewBybitClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	deposites, err := c.GetAllAssetWithNetworks(ctx, &pb.GetAllAssetWithNetworksRequest{})
	var assetList models.CoinInfoResult
	if deposites != nil {
		for _, d := range deposites.Rows {
			var asset models.CoinInfoRow
			asset.Coin = d.Coin
			asset.Name = d.Name
			asset.RemainAmount = d.RemainAmount
			for _, c := range d.Chains {
				var chain models.CoinChainInfo
				chain.Chain = c.Chain
				chain.ChainDeposit = c.ChainDeposit
				chain.ChainType = c.ChainType
				chain.ChainWithdraw = c.ChainWithdraw
				chain.Confirmation = c.Confirmation
				chain.DepositMin = c.DepositMin
				chain.MinAccuracy = c.MinAccuracy
				chain.WithdrawFee = c.WithdrawFee
				chain.WithdrawMin = c.WithdrawMin
				chain.WithdrawPercentageFee = c.WithdrawPercentageFee
				asset.Chains = append(asset.Chains, chain)
			}
			assetList.Rows = append(assetList.Rows, asset)
		}
		if err := mongoDB.Create(ctx, "wallets", "assets", assetList); err != nil {
			log.Println(err)
		}
	}
	log.Println("Ok")
}
