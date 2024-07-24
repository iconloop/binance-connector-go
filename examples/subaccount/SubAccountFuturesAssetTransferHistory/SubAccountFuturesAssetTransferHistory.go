package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	SubAccountFuturesAssetTransferHistory()
}

func SubAccountFuturesAssetTransferHistory() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Query Sub-account Futures Asset Transfer History (For Master Account) - /sapi/v1/sub-account/futures/internalTransfer
	subaccountFuturesAssetTransferHistory, err := client.NewQuerySubAccountFuturesAssetTransferHistoryService().Email("from@email.com").
		FuturesType(1).StartTime(1234567891011).EndTime(1).Page(1).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(subaccountFuturesAssetTransferHistory))
}
