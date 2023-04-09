package models

type Coin struct {
	ID      uint64  `gorm:"primaryKey"`
	CoinID  string  `gorm:"size:255"`
	Name    string  `gorm:"size:255"`
	Symbol  string  `gorm:"size:255"`
	High24h float64 `gorm:"column:orderhy_fundrate;type:decimal(60,20);not null"`
	Low_24h float64 `gorm:"column:orderhy_fundrate;type:decimal(60,20);not null"`
}

func SaveCoin(n string, cId string, s string, h24h float64, l25h float64) *Coin {
	entry := Coin{Name: n, CoinID: cId, Symbol: s, High24h: h24h, Low_24h: l25h}
	DB.Create(&entry)
	return &entry
}

func GetAllCoins() *[]Coin {
	var coins []Coin
	// Get all records
	DB.Find(&coins)
	//result.RowsAffected
	// result.RowsAffected // returns found records count, equals `len(users)`
	// result.Error        // returns error

	return &coins
}
