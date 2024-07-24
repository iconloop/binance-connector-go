package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetDetailOnSubAccountFuturesAccount()
}

func GetDetailOnSubAccountFuturesAccount() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Detail on Sub-account's Futures Account (For Master Account) - /sapi/v1/sub-account/futures/account
	getDetailOnSubAccountFuturesAccount, err := client.NewGetDetailOnSubAccountFuturesAccountService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getDetailOnSubAccountFuturesAccount))
}
