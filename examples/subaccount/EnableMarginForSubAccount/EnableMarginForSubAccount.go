package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	EnableMarginForSubAccount()
}

func EnableMarginForSubAccount() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Enable Margin for Sub-account (For Master Account) - /sapi/v1/sub-account/margin/enable
	enableMarginForSubAccount, err := client.NewEnableMarginForSubAccountService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(enableMarginForSubAccount))
}
