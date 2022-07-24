package main

import (
	"dividend_portfolio_tracker_cli/internal/actions"
	"fmt"
	"os"
)

type Action int64

const (
	CheckPortfolioStatus Action = 1
	PurchaseStock        Action = 2
	SellStock            Action = 3
	ExitApplication      Action = 9
)

func main() {
	fmt.Println("Welcome to dividend portfolio tracker CLI!")

	action := askForAction()

	for {
		switch action {
		case CheckPortfolioStatus:
			fmt.Println("You have selected chec portfolio status action!")
			action = askForAction()
		case PurchaseStock:
			actions.PromptPurchaseStock()
			action = askForAction()
		case SellStock:
			actions.PromptSellStock()
			action = askForAction()
		case ExitApplication:
			os.Exit(0)
		}
	}
}

func askForAction() Action {
	fmt.Print("\n Choose your action (type action number): \n")
	fmt.Printf("%v - Check portfolio status \n", CheckPortfolioStatus)
	fmt.Printf("%v - Purchase stock \n", PurchaseStock)
	fmt.Printf("%v - Sell stock \n", SellStock)
	fmt.Println()
	fmt.Printf("%v - Exit app \n", ExitApplication)
	fmt.Println()

	var action Action
	if _, err := fmt.Scan(&action); err != nil {
		fmt.Printf("Error occured while selecting action: %v", err)
	}

	return action
}
