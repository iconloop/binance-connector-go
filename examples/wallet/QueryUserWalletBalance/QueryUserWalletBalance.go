package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryUserWalletBalance()
}

func QueryUserWalletBalance() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// UserAssetService - /sapi/v3/asset/getUserAsset
	userWalletBalance, err := client.NewQueryUserWalletBalanceService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(userWalletBalance))
}
