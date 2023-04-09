package main

//	"crypto_bot_golang/services"
import (
	"crypto_bot_golang/crypto_services"
	"fmt"
)

func main() {
	c := crypto_services.GetCryptoCoins()
	fmt.Println(c[0].ID)
	coinH := crypto_services.GetCoinPriceHistory("bitcoin")
	fmt.Println(coinH.Prices)
}
