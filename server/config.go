package main

import "time"

const (
	ExchangeRateApiURL        = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	GetExchangeRateApiTimeout = 200 * time.Millisecond
	DatabaseRateInsertTimeout = 10 * time.Millisecond
)
