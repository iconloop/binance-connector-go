package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	DepositAssetsIntoManagedSubAccount()
}

func DepositAssetsIntoManagedSubAccount() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Deposit Assets Into The Managed Sub-account（For Investor Master Account） - /sapi/v1/sub-account/managed-subaccount/deposit
	depositAssetsIntoManagedSubAccount, err := client.NewDepositAssetsIntoManagedSubAccountService().ToEmail("to@email.com").
		Asset("BTC").Amount(0.01).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(depositAssetsIntoManagedSubAccount))
}
