package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	AutoConvertStableCoin()
}

func AutoConvertStableCoin() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// AutoConvertStableCoinService - /sapi/v1/capital/contract/convertible-coins
	autoConvertStableCoin, err := client.NewAutoConvertStableCoinService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(autoConvertStableCoin))
}
