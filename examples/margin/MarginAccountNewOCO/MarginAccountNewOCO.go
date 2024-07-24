package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginAccountNewOCO()
}

func MarginAccountNewOCO() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountNewOCOService - /sapi/v1/margin/order/oco
	marginAccountNewOCO, err := client.NewMarginAccountNewOCOService().Symbol("BTCUSDT").
		Side("BUY").Quantity(0.01).Price(20000).StopPrice(18000).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountNewOCO))
}
