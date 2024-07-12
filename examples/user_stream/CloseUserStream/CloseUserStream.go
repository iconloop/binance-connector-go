package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	CloseUserStream()
}

func CloseUserStream() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	close := client.NewCloseUserStream().ListenKey("your_listen_key").
		Do(context.Background())
	fmt.Println(close)
}
