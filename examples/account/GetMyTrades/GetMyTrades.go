package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetMyTrades()
}

func GetMyTrades() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	symbol := "ICXUSDT"
	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Binance Get trades for a specific account and symbol (USER_DATA) - GET /api/v3/myTrades
	getMyTradesService, err := client.NewGetMyTradesService().
		Symbol(symbol).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getMyTradesService))
}
