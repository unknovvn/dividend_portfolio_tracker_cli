package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const base_client_url = "https://cloud.iexapis.com/stable/stock/"

type DividendData struct {
	Amount      float64 `json:"amount"`
	ExDate      string  `json:"exDate"`
	PaymentDate string  `json:"paymentDate"`
}

type GetQuoteResponse struct {
	CompanyName   string  `json:"companyName"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	LatestPrice   float64 `json:"latestPrice"`
	PeRatio       float64 `json:"peRatio"`
	Week52High    float64 `json:"week52High"`
	Week52Low     float64 `json:"week52Low"`
	YtdChange     float64 `json:"ytdChange"`
}

type GetStockDataResponse struct {
	CompanyName         string
	DayChange           float64
	DayChangePercentage float64
	YtdChange           float64
	LatestPrice         float64
	PeRatio             float64
	Week52High          float64
	Week52Low           float64
	DivAmount           float64
	DivExDate           string
	DivPaymentDate      string
}

func GetStockData(ticker string, token string) (GetStockDataResponse, error) {
	div_data, err := getDividendData(ticker, token)
	if err != nil {
		return GetStockDataResponse{}, err
	}

	quote_data, err := getQuoteData(ticker, token)
	if err != nil {
		return GetStockDataResponse{}, err
	}

	return GetStockDataResponse{
		CompanyName:         quote_data.CompanyName,
		LatestPrice:         quote_data.LatestPrice,
		PeRatio:             quote_data.PeRatio,
		Week52High:          quote_data.Week52High,
		Week52Low:           quote_data.Week52Low,
		DayChange:           quote_data.Change,
		DayChangePercentage: quote_data.ChangePercent,
		YtdChange:           quote_data.YtdChange,
		DivAmount:           div_data.Amount,
		DivExDate:           div_data.ExDate,
		DivPaymentDate:      div_data.PaymentDate,
	}, nil
}

func getDividendData(ticker string, token string) (DividendData, error) {
	div_request_url := base_client_url + ticker + "/dividends/next" + getTokenQueryParam(token)
	div_response, err := http.Get(div_request_url)
	if err != nil {
		fmt.Printf("Error occured while requesting div data: %v\n", err)
		return DividendData{}, err
	}

	defer div_response.Body.Close()

	div_response_body, err := ioutil.ReadAll(div_response.Body)
	if err != nil {
		fmt.Printf("Error occured while requesting div data: %v\n", err)
		return DividendData{}, err
	}

	var response []DividendData
	err = json.Unmarshal(div_response_body, &response)
	if err != nil {
		fmt.Printf("Error occured while requesting div data: %v\n", err)
		return DividendData{}, err
	}

	if len(response) > 0 {
		return response[0], nil
	}

	return DividendData{}, errors.New("No dividend data was returned for ticker: " + ticker)
}

func getQuoteData(ticker string, token string) (GetQuoteResponse, error) {
	quote_request_url := base_client_url + ticker + "/quote" + getTokenQueryParam(token)
	quote_response, err := http.Get(quote_request_url)
	if err != nil {
		return GetQuoteResponse{}, err
	}

	defer quote_response.Body.Close()

	quote_response_body, err := ioutil.ReadAll(quote_response.Body)
	if err != nil {
		return GetQuoteResponse{}, err
	}

	var response GetQuoteResponse
	err = json.Unmarshal(quote_response_body, &response)
	if err != nil {
		return GetQuoteResponse{}, err
	}

	return response, nil
}

func getTokenQueryParam(token string) string {
	return "?token=" + token
}
