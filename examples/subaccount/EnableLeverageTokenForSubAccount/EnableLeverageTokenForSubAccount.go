package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	EnableLeverageTokenForSubAccount()
}

func EnableLeverageTokenForSubAccount() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Enable Leverage Token for Sub-account (For Master Account) - /sapi/v1/sub-account/blvt/enable
	enableLeverageTokenForSubAccount, err := client.NewEnableLeverageTokenForSubAccountService().Email("email@email.com").
		EnableBlvt(true).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(enableLeverageTokenForSubAccount))
}
