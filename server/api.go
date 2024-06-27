package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type ExchangeRateFromUsdToBrl struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func GetExchangeRate() (*ExchangeRateFromUsdToBrl, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, GetExchangeRateApiTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ExchangeRateApiURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var exchangeRate ExchangeRateFromUsdToBrl
	err = json.Unmarshal(body, &exchangeRate)
	if err != nil {
		return nil, err
	}

	return &exchangeRate, nil
}
