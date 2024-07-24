package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetFuturesPositionRiskOfSubAccountV2()
}

func GetFuturesPositionRiskOfSubAccountV2() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Futures Position-Risk of Sub-account V2 (For Master Account) - /sapi/v1/sub-account/futures/positionRisk
	getFuturesPositionRiskOfSubAccountV2, err := client.NewGetFuturesPositionRiskOfSubAccountV2Service().Email("email@email.com").
		FuturesType(1).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getFuturesPositionRiskOfSubAccountV2))
}
