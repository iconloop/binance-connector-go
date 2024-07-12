package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginCurrentOrderCount()
}

func MarginCurrentOrderCount() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginCurrentOrderCountService - /sapi/v1/margin/rateLimit/order
	marginCurrentOrderCount, err := client.NewMarginCurrentOrderCountService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginCurrentOrderCount))
}
