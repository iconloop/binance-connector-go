package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryManagedSubAccountSnapshot()
}

func QueryManagedSubAccountSnapshot() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	withdrawAssetsFromTheManagedSubAccount, err := client.NewQueryManagedSubAccountSnapshotService().Email("email@email.com").
		SubType("BTC").StartTime(123123123).EndTime(123132123).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(withdrawAssetsFromTheManagedSubAccount))
}
