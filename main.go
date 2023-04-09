package main

//	"crypto_bot_golang/services"
import (
	"crypto_bot_golang/services"
	"fmt"
)

func main() {
	x := services.GetCryptoCoins()
	fmt.Println(x)
}
