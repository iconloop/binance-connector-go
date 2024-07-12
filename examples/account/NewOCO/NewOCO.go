package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	NewOCO()
}

func NewOCO() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Binance New OCO (TRADE) - POST /api/v3/order/oco
	newOCO, err := client.NewNewOCOService().Symbol("LTCBNB").
		Side("BUY").Quantity(0.1).Price(0.28).StopPrice(0.22).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(newOCO))
}
