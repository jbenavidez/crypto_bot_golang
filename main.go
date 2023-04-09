package main

//	"crypto_bot_golang/services"
import (
	"crypto_bot_golang/crypto_services"
	"crypto_bot_golang/models"
	"fmt"
)

func main() {
	// init DB
	models.ConnectDB()
	models.DBMigrate()
	// save 3 top coin
	c := crypto_services.SaveTopThreeCoinsProcess()
	fmt.Println(c)
	//coinH := crypto_services.GetCoinPriceHistory("bitcoin")
	//fmt.Println(coinH.Prices)
}
