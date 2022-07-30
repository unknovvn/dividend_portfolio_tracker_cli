package actions

import (
	"bufio"
	"dividend_portfolio_tracker_cli/internal/services"
	"fmt"
	"os"
)

func UpdateApiToken() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Let's update your API token!")
	fmt.Println("You can find your API token here: https://iexcloud.io/cloud-login")
	fmt.Println("What is your new API token value?")

	scanner.Scan()
	api_token := scanner.Text()

	if api_token == "" {
		fmt.Println("Invalid value inserted for API token!")
		return
	}

	services.UpdateApiToken(api_token)
}
