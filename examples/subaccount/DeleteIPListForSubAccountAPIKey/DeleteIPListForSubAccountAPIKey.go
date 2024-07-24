package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	DeleteIPListForSubAccountAPIKey()
}

func DeleteIPListForSubAccountAPIKey() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Delete IP List For a Sub-account API Key (For Master Account) - /sapi/v1/sub-account/subaccountApi/ipRestriction/ipList
	deleteIPListForSubAccountAPIKey, err := client.NewDeleteIPListForSubAccountAPIKeyService().Email("email@email.com").
		SubAccountApiKey("123123").IpAddress("127.0.0.1").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(deleteIPListForSubAccountAPIKey))
}
