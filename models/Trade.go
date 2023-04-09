package models

type Trade struct {
	ID     uint64 `gorm:"primaryKey"`
	CoinID string `gorm:"size:255"`
	Symbol string `gorm:"size:255"`
	// Coin     Coin    `json:"location,omitempty" gorm:"foreignKey:ID;references:ID"`
	Quantity int     `gorm:"size:255"`
	cost     float64 `gorm:"size:255"`
	AvgCost  float64 `gorm:"size:255"`
}

type CoinTobuy struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	CurrentPrice float64 `json:"current_price"`
	IsValidToBuy bool    `json:"is_valid"`
	AvgPrice     float64 `json:"avg_price"`
}

func PlaceOrder(c CoinTobuy, q int) *Trade {
	entry := Trade{
		CoinID:   c.ID,
		Symbol:   c.Symbol,
		Quantity: q,
		cost:     c.CurrentPrice,
		AvgCost:  c.AvgPrice,
	}
	return &entry
}
