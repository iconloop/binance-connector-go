package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetManagedSubAccountDepositAddress()
}

func GetManagedSubAccountDepositAddress() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Managed Sub-account Deposit Address (For Investor Master Account) - /sapi/v1/managed-subaccount/deposit/address
	GetManagedSubAccountDepositAddress, err := client.NewGetManagedSubAccountDepositAddressService().Email("from@email.com").
		Coin("BTC").Network("BTC").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(GetManagedSubAccountDepositAddress))
}
