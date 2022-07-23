package internal

type StockData struct {
	Ticker       string  `json:"ticker"`
	Shares       int     `json:"shares"`
	Price        float64 `json:"price"`
	PurchaseDate int64   `json:"purchaseDate"`
}

type PortfolioData struct {
	Stocks []StockData `json:"stocks"`
}
