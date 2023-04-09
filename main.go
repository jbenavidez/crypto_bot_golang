package main

//	"crypto_bot_golang/services"
import (
	"crypto_bot_golang/crypto_services"
	"fmt"
)

func main() {
	x := crypto_services.GetCryptoCoins()
	fmt.Println(x)
}
