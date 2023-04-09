package crypto_services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
)

var client *http.Client

type Coins []struct {
	ID                           string    `json:"id"`
	Symbol                       string    `json:"symbol"`
	Name                         string    `json:"name"`
	Image                        string    `json:"image"`
	CurrentPrice                 float64   `json:"current_price"`
	MarketCap                    float64   `json:"market_cap"`
	MarketCapRank                float64   `json:"market_cap_rank"`
	FullyDilutedValuation        float64   `json:"fully_diluted_valuation"`
	TotalVolume                  float64   `json:"total_volume"`
	High24H                      float64   `json:"high_24h"`
	Low24H                       float64   `json:"low_24h"`
	PriceChange24H               float64   `json:"price_change_24h"`
	PriceChangePercentage24H     float64   `json:"price_change_percentage_24h"`
	MarketCapChange24H           float64   `json:"market_cap_change_24h"`
	MarketCapChangePercentage24H float64   `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            float64   `json:"circulating_supply"`
	TotalSupply                  float64   `json:"total_supply"`
	MaxSupply                    float64   `json:"max_supply"`
	Ath                          float64   `json:"ath"`
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

// This function will get the top 3 coins at the current time,
func GetTopThreeCoins() Coins {
	client = &http.Client{Timeout: 10 * time.Second}
	coins := GetCoins()
	//sort to get the top 3 coins
	sort.Slice(coins[:], func(i, j int) bool {
		return coins[i].CurrentPrice > coins[j].CurrentPrice
	})
	return coins[0:3]
}

// Get the coins at the current time.
func GetCoins() Coins {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10&page=1&sparkline=false"
	var coins Coins
	err := GetJson(url, &coins)
	if err != nil {
		fmt.Println("error getting coins", err)
	} else {
		fmt.Println("coins retrieved", len(coins))
	}
	// Important keys
	// - id
	// - symbol
	// - name
	// - current_price

	return coins

}

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
	// Important keys
	// Item 0 -> Unix Timestamp
	// Item 1 -> price
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
