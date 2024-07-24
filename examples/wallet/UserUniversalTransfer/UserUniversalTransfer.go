package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	UserUniversalTransfer()
}

func UserUniversalTransfer() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// UserUniversalTransferService - /sapi/v1/asset/transfer
	userUniversalTransfer, err := client.NewUserUniversalTransferService().
		TransferType("MAIN_UMFUTURE").Asset("USDT").Amount(20.50).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(userUniversalTransfer))
}
