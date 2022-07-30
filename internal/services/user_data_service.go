package services

import (
	"dividend_portfolio_tracker_cli/internal"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const user_data_file_path = ".dividend_portfolio_tracker"

func PurchaseStock(ticker string, shares float64, price float64, date time.Time) {
	newStock := internal.TransactionData{
		Shares:       shares,
		Price:        price,
		PurchaseDate: date.Unix(),
		Operation:    internal.PurchaseOperation,
	}

	insertStockOperation(ticker, newStock)
}

func SellStock(ticker string, shares float64, price float64, date time.Time) {
	newStock := internal.TransactionData{
		Shares:       shares,
		Price:        price,
		PurchaseDate: date.Unix(),
		Operation:    internal.SellOperation,
	}

	insertStockOperation(ticker, newStock)
}

func UpdateApiToken(apiToken string) {
	user_data := GetUserData()
	user_data.ApiToken = apiToken

	err := saveUserData(user_data)
	if err != nil {
		fmt.Printf("Error occured while saving user data: %v", err)
	}
}

func GetUserData() internal.UserData {
	path, err := getUserDataFilePath()
	if err != nil {
		return internal.UserData{
			Transactions: make(map[string][]internal.TransactionData),
		}
	}

	user_data_content, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		if _, err := os.Create(path); err != nil {
			fmt.Printf("Error occured while creating data.json file: %v", err)
		}

		user_data := internal.UserData{
			Transactions: make(map[string][]internal.TransactionData),
		}

		saveUserData(user_data)

		return user_data
	} else {
		var user_data internal.UserData
		if err := json.Unmarshal(user_data_content, &user_data); err != nil {
			fmt.Printf("Unable to unmarshal user data file: %v", err)

			return internal.UserData{
				Transactions: make(map[string][]internal.TransactionData),
			}
		}

		return user_data
	}
}

func insertStockOperation(ticker string, newStock internal.TransactionData) {
	user_data := GetUserData()

	if stocks, ok := user_data.Transactions[ticker]; ok {
		user_data.Transactions[ticker] = append(stocks, newStock)
	} else {
		user_data.Transactions[ticker] = []internal.TransactionData{newStock}
	}

	err := saveUserData(user_data)
	if err != nil {
		fmt.Printf("Error occured while saving user data: %v", err)
	}
}

func saveUserData(userData internal.UserData) error {
	path, err := getUserDataFilePath()
	if err != nil {
		return err
	}

	user_data_json, err := json.Marshal(userData)
	if err != nil {
		return err
	}

	return os.WriteFile(path, user_data_json, os.ModeAppend)
}

func getUserDataFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + string(os.PathSeparator) + user_data_file_path, nil
}
