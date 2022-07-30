package internal

type TransactionData struct {
	Shares       float64   `json:"shares"`
	Price        float64   `json:"price"`
	PurchaseDate int64     `json:"purchaseDate"`
	Operation    Operation `json:"operation"`
}

type UserData struct {
	ApiToken     string                       `json:"api_token"`
	Transactions map[string][]TransactionData `json:"stocks"`
}

type Operation int

const (
	PurchaseOperation = 1
	SellOperation     = 2
)
