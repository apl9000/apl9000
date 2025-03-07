package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ExchangeRateResponse struct {
	Result             string             `json:"result"`
	Provider           string             `json:"provider"`
	Documentation      string             `json:"documentation"`
	TermsOfUse         string             `json:"terms_of_use"`
	TimeLastUpdateUnix int64              `json:"time_last_update_unix"`
	TimeLastUpdateUTC  string             `json:"time_last_update_utc"`
	TimeNextUpdateUnix int64              `json:"time_next_update_unix"`
	TimeNextUpdateUTC  string             `json:"time_next_update_utc"`
	TimeEOLUnix        int64              `json:"time_eol_unix"`
	BaseCode           string             `json:"base_code"`
	Rates              map[string]float64 `json:"rates"`
}

type Rates struct {
	USD string
	CAD string
	EUR string
	GBP string
	JPY string
	MXN string
	CNY string
	INR string
	AUD string
	ZAR string
	RUB string
	BRL string
	CHF string
}

func GetRates() Rates {
	// https://www.exchangerate-api.com/docs/free
	uri := "https://open.er-api.com/v6/latest/CAD"

	response, err := http.Get(uri)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return Rates{}
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return Rates{}
	}
	var data ExchangeRateResponse
	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
		return Rates{}
	}

	return Rates{
		CAD: fmt.Sprintf("%.2f", data.Rates["CAD"]),
		USD: fmt.Sprintf("%.2f", data.Rates["USD"]),
		EUR: fmt.Sprintf("%.2f", data.Rates["EUR"]),
		GBP: fmt.Sprintf("%.2f", data.Rates["GBP"]),
		JPY: fmt.Sprintf("%.2f", data.Rates["JPY"]),
		MXN: fmt.Sprintf("%.2f", data.Rates["MXN"]),
		AUD: fmt.Sprintf("%.2f", data.Rates["AUD"]),
	}
}
