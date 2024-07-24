package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	CrossMarginAccountDetail()
}

func CrossMarginAccountDetail() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// CrossMarginAccountDetailService - /sapi/v1/margin/account
	crossMarginAccountDetail, err := client.NewCrossMarginAccountDetailService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(crossMarginAccountDetail))
}
