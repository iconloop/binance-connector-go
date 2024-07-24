package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QuerySubAccountSpotAssetsSummary()
}

func QuerySubAccountSpotAssetsSummary() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Query Sub-account Spot Assets Summary (For Master Account) - /sapi/v1/sub-account/spotSummary
	querySubAccountSpotAssetsSummary, err := client.NewQuerySubAccountSpotAssetsSummaryService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(querySubAccountSpotAssetsSummary))
}
