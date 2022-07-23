package services

import (
	"dividend_portfolio_tracker_cli/internal"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const portfolio_data_file_path = ".dividend_portfolio_tracker"

func PurchaseStock(ticker string, shares int, price float64, date time.Time) {
	newStock := internal.StockData{
		Ticker:       ticker,
		Shares:       shares,
		Price:        price,
		PurchaseDate: date.Unix(),
	}

	portfolio_data := GetPortfolioData()
	portfolio_data.Stocks = append(portfolio_data.Stocks, newStock)

	err := savePortfolioData(portfolio_data)
	if err != nil {
		fmt.Printf("Error occured while saving portfolio data: %v", err)
	}
}

func GetPortfolioData() internal.PortfolioData {
	path, err := getPortfolioDataFilePath()
	if err != nil {
		return internal.PortfolioData{
			Stocks: []internal.StockData{},
		}
	}

	portfolio_data_content, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		if _, err := os.Create(path); err != nil {
			fmt.Printf("Error occured while creating data.json file: %v", err)
		}

		return internal.PortfolioData{
			Stocks: []internal.StockData{},
		}
	} else {
		var portfolio_data internal.PortfolioData
		if err := json.Unmarshal(portfolio_data_content, &portfolio_data); err != nil {
			fmt.Printf("Unable to unmarshal user data file: %v", err)
		}

		return portfolio_data
	}
}

func savePortfolioData(portfolioData internal.PortfolioData) error {
	path, err := getPortfolioDataFilePath()
	if err != nil {
		return err
	}

	portfolio_data_json, err := json.Marshal(portfolioData)
	if err != nil {
		return err
	}

	return os.WriteFile(path, portfolio_data_json, os.ModeAppend)
}

func getPortfolioDataFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + string(os.PathSeparator) + portfolio_data_file_path, nil
}
