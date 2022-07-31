package actions

import (
	"dividend_portfolio_tracker_cli/internal"
	"dividend_portfolio_tracker_cli/internal/clients"
	"dividend_portfolio_tracker_cli/internal/services"
	"fmt"
	"sort"

	"github.com/alexeyco/simpletable"
)

func CheckPortfolioStatus() {
	portfolio_data := services.GetUserData()
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
	fmt.Println()
	printSummary(stock_statuses)
}

func addTableHeaders(table *simpletable.Table) {
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Company name"},
			{Align: simpletable.AlignCenter, Text: "Shares"},
			{Align: simpletable.AlignCenter, Text: "Buy Price"},
			{Align: simpletable.AlignLeft, Text: "Market Price"},
			{Align: simpletable.AlignLeft, Text: "Buy Value"},
			{Align: simpletable.AlignLeft, Text: "Market Value"},
			{Align: simpletable.AlignLeft, Text: "Pe"},
			{Align: simpletable.AlignLeft, Text: "Day Change"},
			{Align: simpletable.AlignLeft, Text: "Day Change %"},
			{Align: simpletable.AlignLeft, Text: "Ytd Change"},
			{Align: simpletable.AlignLeft, Text: "52W High"},
			{Align: simpletable.AlignLeft, Text: "52W Low"},
			{Align: simpletable.AlignLeft, Text: "Annual Div"},
			{Align: simpletable.AlignLeft, Text: "Div Yield"},
			{Align: simpletable.AlignLeft, Text: "Div Yield on Cost"},
			{Align: simpletable.AlignLeft, Text: "Div Ex Date"},
			{Align: simpletable.AlignLeft, Text: "Div Payment Date"},
		},
	}
}

func addTableBody(table *simpletable.Table, stock_statuses []StockStatus) {
	sort.Slice(stock_statuses, func(i, j int) bool {
		return stock_statuses[i].MarketValue > stock_statuses[j].MarketValue
	})

	for _, stock_status := range stock_statuses {
		if stock_status.Shares <= 0 {
			continue
		}

		row := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: stock_status.Name},
			{Align: simpletable.AlignLeft, Text: fmt.Sprint(stock_status.Shares)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f $", stock_status.BuyPrice)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f $", stock_status.MarketPrice)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f $", stock_status.BuyValue)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f $", stock_status.MarketValue)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f $", stock_status.PeRatio)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f $", stock_status.DayChange)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f %%", stock_status.DayChangePercentage)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f %%", stock_status.YtdChange)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f $", stock_status.Week52High)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f $", stock_status.Week52Low)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f $", stock_status.DivAnnual)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f %%", stock_status.DivYield)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f %%", stock_status.DivYOC)},
			{Align: simpletable.AlignLeft, Text: stock_status.DivExDate},
			{Align: simpletable.AlignLeft, Text: stock_status.DivPaymentDate},
		}

		table.Body.Cells = append(table.Body.Cells, row)
	}
}

func printSummary(stock_statuses []StockStatus) {
	totalBuyValue := 0.0
	for _, stock_status := range stock_statuses {
		totalBuyValue += stock_status.BuyValue
	}

	fmt.Printf("Total portfolio buy value: %v\n", fmt.Sprintf("%.2f $", totalBuyValue))
}

func getStockStatuses(portfolio_data internal.UserData) []StockStatus {
	stockStatuses := []StockStatus{}
	for ticker, transactions := range portfolio_data.Transactions {

		shares := 0.0
		buy_value := 0.0
		for _, transaction := range transactions {
			if transaction.Operation == internal.PurchaseOperation {
				shares += transaction.Shares
				buy_value += transaction.Price * transaction.Shares
			} else {
				shares -= transaction.Shares
				buy_value -= transaction.Price * transaction.Shares
			}
		}

		newStockStatus := StockStatus{
			Name:     ticker,
			Shares:   shares,
			BuyValue: buy_value,
			BuyPrice: buy_value / shares,
		}

		if portfolio_data.ApiToken != "" {
			stock_data, err := clients.GetStockData(ticker, portfolio_data.ApiToken)
			if err != nil {
				fmt.Printf("Error returned from GetStockData: %v\n", err)
			} else {
				newStockStatus.Name = stock_data.CompanyName
				newStockStatus.MarketPrice = stock_data.LatestPrice
				newStockStatus.MarketValue = stock_data.LatestPrice * shares
				newStockStatus.DayChange = stock_data.DayChange
				newStockStatus.DayChangePercentage = stock_data.DayChangePercentage * 100
				newStockStatus.YtdChange = stock_data.YtdChange * 100
				newStockStatus.PeRatio = stock_data.PeRatio
				newStockStatus.Week52High = stock_data.Week52High
				newStockStatus.Week52Low = stock_data.Week52Low
				newStockStatus.DivAnnual = stock_data.DivAnnual * shares
				newStockStatus.DivYield = stock_data.DivYield
				newStockStatus.DivYOC = stock_data.DivAnnual * shares / buy_value * 100
				newStockStatus.DivExDate = stock_data.DivExDate
				newStockStatus.DivPaymentDate = stock_data.DivPaymentDate
			}
		}

		stockStatuses = append(stockStatuses, newStockStatus)
	}

	return stockStatuses
}

type StockStatus struct {
	Name                string
	Shares              float64
	BuyPrice            float64
	MarketPrice         float64
	BuyValue            float64
	MarketValue         float64
	PeRatio             float64
	DayChange           float64
	DayChangePercentage float64
	YtdChange           float64
	Week52High          float64
	Week52Low           float64
	DivAnnual           float64
	DivYield            float64
	DivYOC              float64
	DivExDate           string
	DivPaymentDate      string
}
