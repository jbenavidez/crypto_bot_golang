package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

type Coin struct {
	Name                     string  `json:"name"`
	Symbol                   string  `json:"symbol"`
	Image                    string  `json:"image"`
	CurrentPrice             float64 `json:"current_price"`
	MarkCap                  int     `json:"mark_cap"`
	MarketCapRank            int     `json:"market_cap_rank"`
	FullyDilutedValuation    int     `json:"fully_diluted_valuation"`
	TotalVolume              int     `json:"total_volume"`
	High24h                  float64 `json:"high_24h"`
	Low24h                   float64 `json:"low_24h"`
	PriceChange24h           float64 `json:"price_change_24h"`
	PriceChangePercentage24h float64 `json:"price_change_percentage_24h"`
	MarketCapChange24h       float64 `json:"market_cap_change_24h"`
	CirculatingSupply        float64 `json:"circulating_supply"`
	TotalSupply              float64 `json:"total_supply"`
	MaxSupply                float64 `json:"max_supply"`
	Ath                      float64 `json:"ath"`
	AthChangePercentage      float64 `json:"ath_change_percentage"`
	AthDate                  string  `json:"ath_date"`
	Atl                      float64 `json:"atl"`
	AtlChangePercentage      float64 `json:"atl_change_percentage"`
	AtlDate                  string  `json:"atl_date"`
	LastUpdated              string  `json:"last_updated"`
}

type Coins []Coin

func GetCryptoCoins() Coins {
	client = &http.Client{Timeout: 10 * time.Second}
	coins := TopCoins()
	return coins
}

func TopCoins() Coins {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10&page=1&sparkline=false"
	var coins Coins
	err := GetJson(url, &coins)
	if err != nil {
		fmt.Println("error getting coins", err)
	} else {
		fmt.Println("coins retrieved", len(coins))
	}
	return coins

}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
