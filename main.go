package main

//	"crypto_bot_golang/services"
import (
	"crypto_bot_golang/crypto_services"
	"fmt"
)

func main() {
	c := crypto_services.GetTopThreeCoins()

	for _, v := range c {
		fmt.Println(v.CurrentPrice)
	}
	//coinH := crypto_services.GetCoinPriceHistory("bitcoin")
	//fmt.Println(coinH.Prices)
}
