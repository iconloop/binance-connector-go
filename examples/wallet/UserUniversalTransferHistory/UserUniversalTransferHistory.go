package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	UserUniversalTransferHistory()
}

func UserUniversalTransferHistory() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// UserUniversalTransferHistoryService - /sapi/v1/asset/transfer
	userUniversalTransferHistory, err := client.NewUserUniversalTransferHistoryService().
		TransferType("MAIN_UMFUTURE").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(userUniversalTransferHistory))
}
