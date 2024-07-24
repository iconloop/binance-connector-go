package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WithdrawAssetsFromTheManagedSubAccount()
}

func WithdrawAssetsFromTheManagedSubAccount() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	withdrawAssetsFromTheManagedSubAccount, err := client.NewWithdrawAssetsFromTheManagedSubAccountService().FromEmail("email@email.com").
		Asset("BTC").Amount(1.5).TransferDate(123132123).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(withdrawAssetsFromTheManagedSubAccount))
}
