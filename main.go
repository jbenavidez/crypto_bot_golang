package main

//	"crypto_bot_golang/services"
import (
	"crypto_bot_golang/cryptoServices"
	"crypto_bot_golang/models"
)

func main() {
	// init DB
	models.ConnectDB()
	models.DBMigrate()
	cryptoServices.SaveTopThreeCoinsProcess() // save 3 top coin
	cryptoServices.ComparePlaceCoinsPrice()   // buy coins

}
