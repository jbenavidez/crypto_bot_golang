package crypto_services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

type Coins []struct {
	ID                           string    `json:"id"`
	Symbol                       string    `json:"symbol"`
	Name                         string    `json:"name"`
	Image                        string    `json:"image"`
	CurrentPrice                 int       `json:"current_price"`
	MarketCap                    int64     `json:"market_cap"`
	MarketCapRank                int       `json:"market_cap_rank"`
	FullyDilutedValuation        int64     `json:"fully_diluted_valuation"`
	TotalVolume                  int64     `json:"total_volume"`
	High24H                      int       `json:"high_24h"`
	Low24H                       int       `json:"low_24h"`
	PriceChange24H               float64   `json:"price_change_24h"`
	PriceChangePercentage24H     float64   `json:"price_change_percentage_24h"`
	MarketCapChange24H           int       `json:"market_cap_change_24h"`
	MarketCapChangePercentage24H float64   `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            float64   `json:"circulating_supply"`
	TotalSupply                  float64   `json:"total_supply"`
	MaxSupply                    float64   `json:"max_supply"`
	Ath                          int       `json:"ath"`
	AthChangePercentage          float64   `json:"ath_change_percentage"`
	AthDate                      time.Time `json:"ath_date"`
	Atl                          float64   `json:"atl"`
	AtlChangePercentage          float64   `json:"atl_change_percentage"`
	AtlDate                      time.Time `json:"atl_date"`
	Roi                          any       `json:"roi"`
	LastUpdated                  time.Time `json:"last_updated"`
}
type Prices struct {
	Prices       [][]float64 `json:"prices"`
	MarketCaps   [][]float64 `json:"market_caps"`
	TotalVolumes [][]float64 `json:"total_volumes"`
}

// This function will get the top 10 coins at the current time,
// sorted by market cap in desc order.
func GetCryptoCoins() Coins {
	client = &http.Client{Timeout: 10 * time.Second}
	coins := TopCoins()
	return coins
}

// Important keys
// - id
// - symbol
// - name
// - current_price
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

// Returns a slide of slide
// Item 0 -> Unix Timestamp
// Item 1 -> price
func GetCoinPriceHistory(coinId string) Prices {
	client = &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=usd&days=9&interval=daily", coinId)
	var prices Prices
	err := GetJson(url, &prices)
	if err != nil {
		fmt.Println("error getting coins", err)
	} else {
		fmt.Println("Coin prices retrieved")
	}
	return prices
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
