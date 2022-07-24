package actions

import (
	"dividend_portfolio_tracker_cli/internal"
	"dividend_portfolio_tracker_cli/internal/services"
	"fmt"

	"github.com/alexeyco/simpletable"
)

func CheckPortfolioStatus() {
	portfolio_data := services.GetPortfolioData()
	stock_statuses := getStockStatuses(portfolio_data)

	if len(stock_statuses) == 0 {
		fmt.Println("There is no data to be displayed! Make stock purchases first ;)")
		return
	}

	table := simpletable.New()
	addTableHeaders(table)
	addTableBody(table, stock_statuses)

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func addTableHeaders(table *simpletable.Table) {
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Company name"},
			{Align: simpletable.AlignCenter, Text: "Shares"},
			{Align: simpletable.AlignCenter, Text: "Buy Price"},
			{Align: simpletable.AlignLeft, Text: "Buy Value"},
		},
	}
}

func addTableBody(table *simpletable.Table, stock_statuses []StockStatus) {
	for _, stock_status := range stock_statuses {
		if stock_status.Shares <= 0 {
			continue
		}

		row := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: stock_status.Name},
			{Align: simpletable.AlignLeft, Text: fmt.Sprint(stock_status.Shares)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f $", stock_status.BuyPrice)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f $", stock_status.BuyValue)},
		}

		table.Body.Cells = append(table.Body.Cells, row)
	}
}

func getStockStatuses(portfolio_data internal.PortfolioData) []StockStatus {

	stockStatuses := []StockStatus{}
	for ticker, transactions := range portfolio_data.Transactions {

		shares := 0
		buyValue := 0.0
		for _, transaction := range transactions {
			if transaction.Operation == internal.PurchaseOperation {
				shares += transaction.Shares
				buyValue += transaction.Price * float64(transaction.Shares)
			} else {
				shares -= transaction.Shares
				buyValue -= transaction.Price * float64(transaction.Shares)
			}
		}

		newStockStatus := StockStatus{
			Name:     ticker, // todo: change to company name
			Shares:   shares,
			BuyValue: buyValue,
			BuyPrice: buyValue / float64(shares),
		}

		stockStatuses = append(stockStatuses, newStockStatus)
	}

	return stockStatuses
}

type StockStatus struct {
	Name     string
	Shares   int
	BuyPrice float64
	BuyValue float64
}
