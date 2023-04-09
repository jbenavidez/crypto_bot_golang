package cryptoServices

import (
	"crypto_bot_golang/models"
	"fmt"
	"log"
	"sort"
)

// # Get top 3 coins
func GetTopThreeCoins() Coins {
	coins := GetCoins()
	//sort to get the top 3 coins
	sort.Slice(coins[:], func(i, j int) bool {
		return coins[i].CurrentPrice > coins[j].CurrentPrice
	})
	return coins[0:3]
}
func insertOnDb(c Coins) {
	//
	log.Println("init saving coins to db")
	for _, v := range c {
		log.Println("saving coins to db", v.Name)
		models.SaveCoin(v.Name, v.ID, v.Symbol, v.High24H, v.Low24H)
		log.Printf("%s was successfully Sucefully", v.Name)
	}

}

func SaveTopThreeCoinsProcess() string {
	// c := GetTopThreeCoins()
	// insertOnDb(c)
	notes := models.GetAllCoins()
	fmt.Println("the nores", notes)

	return "hello2"
}
