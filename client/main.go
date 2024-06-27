package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ExchangeRateFromUsdToBrl struct {
	Bid string `json:"bid"`
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotation", nil)
	CheckError(err)

	res, err := http.DefaultClient.Do(req)
	CheckError(err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Status code is not OK")
		return
	}

	body, err := io.ReadAll(res.Body)
	CheckError(err)

	var exchangeRate ExchangeRateFromUsdToBrl
	fmt.Println(string(body))
	err = json.Unmarshal(body, &exchangeRate)
	CheckError(err)

	file, err := os.Create("cotacao.txt")
	CheckError(err)
	defer file.Close()

	_, err = file.WriteString("Dólar: " + exchangeRate.Bid)
	CheckError(err)
}
