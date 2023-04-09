package cryptoServices

import (
	"crypto_bot_golang/models"
	"fmt"
	"log"

	"github.com/montanaflynn/stats"
)

// Calculate avg price for a given coin
func CalAvgPrice(c Coin) models.CoinTobuy {
	log.Println("calculating avg price for ", c.Name)
	ph := GetCoinPriceHistory(c.ID)
	var coinPrice []float64
	for _, v := range ph.Prices {
		coinPrice = append(coinPrice, v[1])
	}
	avg, err := stats.Mean(coinPrice) //avg price
	if err != nil {
		fmt.Println(err)
	}
	coin := models.CoinTobuy{
		ID:           c.ID,
		Symbol:       c.Symbol,
		CurrentPrice: c.CurrentPrice,
		IsValidToBuy: false,
		AvgPrice:     avg,
	}
	// check if place order is needed
	log.Println("checking price for ", coin.ID)
	if c.CurrentPrice < avg {
		potentialGain := avg - c.CurrentPrice
		coin.IsValidToBuy = true
		log.Printf("Potential gain for %v\n:   %v\n", coin.ID, potentialGain)
	}

	return coin
}

// if avg price is below current, place order
func ComparePlaceCoinsPrice() {
	c := GetCoins()
	for _, v := range c {
		coin := CalAvgPrice(v)
		if coin.IsValidToBuy {
			log.Printf("placing order for: ", coin.ID)
			models.PlaceOrder(coin, 1)
			log.Println("Order created  for: ", coin.ID)
		}

	}

}
