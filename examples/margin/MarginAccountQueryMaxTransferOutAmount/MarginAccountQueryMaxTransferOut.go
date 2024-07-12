package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginAccountQueryMaxTransferOutAmount()
}

func MarginAccountQueryMaxTransferOutAmount() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountQueryMaxTransferOutAmountService - /sapi/v1/margin/maxTransferable
	marginAccountQueryMaxTransferOutAmount, err := client.NewMarginAccountQueryMaxTransferOutAmountService().
		Asset("USDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountQueryMaxTransferOutAmount))
}
