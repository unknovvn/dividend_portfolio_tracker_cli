package actions

import (
	"bufio"
	"dividend_portfolio_tracker_cli/internal/services"
	"fmt"
	"os"
	"strconv"
	"time"
)

func PromptPurchaseStock() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Let's add new stock position to your portfolio!")
	fmt.Println("What is the stock ticker?")
	scanner.Scan()
	ticker := scanner.Text()
	fmt.Println("What is the amount of stocks bought?")
	scanner.Scan()
	amount_string := scanner.Text()
	amount, err := strconv.ParseFloat(amount_string, 64)
	if err != nil {
		fmt.Println("Invalid value inserted for amount")
		return
	}

	fmt.Println("What is the cost per share?")
	scanner.Scan()
	cost_per_share_string := scanner.Text()
	cost_per_share, err := strconv.ParseFloat(cost_per_share_string, 64)
	if err != nil {
		fmt.Println("Invalid value inserted for cost per share")
		return
	}

	fmt.Println("Optional: What is the date of transaction?")
	scanner.Scan()
	transaction_date_string := scanner.Text()

	var transaction_date time.Time
	if transaction_date_string == "" {
		transaction_date = time.Now()
	} else {
		transaction_date, err = time.Parse(time.RFC3339, transaction_date_string)
		if err != nil {
			fmt.Println("Invalid value inserted for date of transaction")
			return
		}
	}

	services.PurchaseStock(ticker, amount, cost_per_share, transaction_date)
}
