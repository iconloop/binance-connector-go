package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginSmallLiabilityExchangeHistory()
}

func MarginSmallLiabilityExchangeHistory() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginSmallLiabilityExchangeHistoryService - /sapi/v1/margin/exchange-small-liability-history
	marginSmallLiabilityExchangeHistory, err := client.NewMarginSmallLiabilityExchangeHistoryService().
		Current(1).Size(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginSmallLiabilityExchangeHistory))
}
