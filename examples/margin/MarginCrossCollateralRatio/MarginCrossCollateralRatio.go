package main

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginCrossCollateralRatio()
}

func MarginCrossCollateralRatio() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginCrossCollateralRatioService - /sapi/v1/margin/crossMarginCollateralRatio
	marginCrossCollateralRatio, err := client.NewMarginCrossCollateralRatioService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginCrossCollateralRatio))
}
