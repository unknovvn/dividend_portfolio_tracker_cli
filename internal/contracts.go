package internal

type StockData struct {
	Shares       int       `json:"shares"`
	Price        float64   `json:"price"`
	PurchaseDate int64     `json:"purchaseDate"`
	Operation    Operation `json:"operation"`
}

type PortfolioData struct {
	Stocks map[string][]StockData `json:"stocks"`
}

type Operation int

const (
	PurchaseOperation = 1
	SellOperation     = 2
)
