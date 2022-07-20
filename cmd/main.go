package main

import (
	"fmt"
	"os"
)

type Action int64

const (
	CheckPortfolioStatus Action = 1
)

func main() {
	fmt.Println("Welcome to dividend portfolio tracker CLI!")

	action := askForAction()

	for {
		switch action {
		case CheckPortfolioStatus:
			fmt.Println("You have selected chec portfolio status action!")
			action = askForAction()
		case 9:
			os.Exit(0)
		}
	}
}

func askForAction() Action {
	fmt.Print("\nChoose your action (type action number):\n")
	fmt.Printf("%v - Check portfolio status \n", CheckPortfolioStatus)
	fmt.Println()
	fmt.Println("9 - Exit app")
	fmt.Println()

	var action Action
	if _, err := fmt.Scan(&action); err != nil {
		fmt.Printf("Error occured while selecting action: %v", err)
	}

	return action
}
