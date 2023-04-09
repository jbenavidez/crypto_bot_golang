package cryptoServices

import (
	"fmt"
	"log"

	"github.com/montanaflynn/stats"
)

type CoinTobuy struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	CurrentPrice float64 `json:"current_price"`
	IsValidToBuy bool    `json:"is_valid"`
	AvgPrice     float64 `json:"avg_price"`
}

// Calculate avg price for a given coin
func CalAvgPrice(c Coin) Coin {
	//log.Println("calculating avg price for %s", c.Name)
	ph := GetCoinPriceHistory(c.ID)
	var coinPrice []float64
	for _, v := range ph.Prices {
		coinPrice = append(coinPrice, v[1])
	}
	avg, err := stats.Mean(coinPrice) //avg price
	if err != nil {
		fmt.Println(err)
	}
	coin := CoinTobuy{
		ID:           c.ID,
		Symbol:       c.Symbol,
		CurrentPrice: c.CurrentPrice,
		IsValidToBuy: false,
		AvgPrice:     avg,
	}
	// check if place order is needed
	log.Println("checking price for %s", coin.ID)
	if c.CurrentPrice < avg {
		potentialGain := avg - c.CurrentPrice
		coin.IsValidToBuy = true
		log.Println("Potential gain for %s: %%", coin.ID, potentialGain)
	}

	return c
}

func ComparePlaceCoinsPrice() {
	c := GetCoins()
	for _, v := range c {
		CalAvgPrice(v)
	}

}
