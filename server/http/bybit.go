package http

import (
	"context"
	"encoding/json"
	"fmt"

	bybit_connector "github.com/wuhewuhe/bybit.go.api"
	"github.com/wuhewuhe/bybit.go.api/models"
)

type BybitHttp struct {
	client *bybit_connector.Client
}

func NewBybitHttp(apiKey, apiSecretkey string) *BybitHttp {
	client := bybit_connector.NewBybitHttpClient(apiKey, apiSecretkey, bybit_connector.WithBaseURL(bybit_connector.MAINNET))
	return &BybitHttp{client}
}

func (b *BybitHttp) GetCoinInfo() (error, models.CoinInfoResult) {
	params := map[string]interface{}{"": ""}
	assetResult, err := b.client.NewUtaBybitServiceWithParams(params).GetCoinInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return err, models.CoinInfoResult{}
	}
	var result models.CoinInfoResult
	json.Unmarshal([]byte(bybit_connector.PrettyPrint(assetResult.Result)), &result)
	return err, result
}
