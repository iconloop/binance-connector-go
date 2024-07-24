package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	SubAccountTransferHistory()
}

func SubAccountTransferHistory() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Sub-account Transfer History (For Sub-account) - /sapi/v1/sub-account/transfer/subUserHistory
	subAccountTransferHistory, err := client.NewSubAccountTransferHistoryService().Asset("BTC").
		TransferType(1).StartTime(1234567891011).EndTime(1234567891011).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(subAccountTransferHistory))
}
