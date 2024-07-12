package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryManagedSubAccountAssetDetails()
}

func QueryManagedSubAccountAssetDetails() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Query Managed Sub-account Asset Details（For Investor Master Account）- /sapi/v1/sub-account/managed-subaccount/asset
	queryManagedSubAccountAssetDetails, err := client.NewQueryManagedSubAccountAssetDetailsService().Email("email@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryManagedSubAccountAssetDetails))
}
