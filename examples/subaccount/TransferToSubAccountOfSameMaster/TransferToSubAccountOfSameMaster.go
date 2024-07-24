package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	TransferToSubAccountOfSameMaster()
}

func TransferToSubAccountOfSameMaster() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Transfer to Sub-account of Same Master (For Sub-account) - /sapi/v1/sub-account/transfer/subToSub
	transferToSubAccountOfSameMaster, err := client.NewTransferToSubAccountOfSameMasterService().ToEmail("from@email.com").Asset("BTC").
		Amount(0.01).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(transferToSubAccountOfSameMaster))
}
