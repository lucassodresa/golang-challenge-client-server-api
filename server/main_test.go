package main

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

// happy path
func TestShouldRespondWithExchangeRateWhenExchangeApiReturnsOk(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", ExchangeRateApiURL,
		httpmock.NewStringResponder(200, `{"USDBRL": {"code": "USD", "codein": "BRL", "name": "Dólar Americano/Real Brasileiro", "high": "5.30", "low": "5.30", "varBid": "0.00", "pctChange": "0.00", "bid": "5.30", "ask": "5.30", "timestamp": "1622030400", "create_date": "2021-05-26 17:00:01"}}`))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	CotationHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200 but got %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expected := `{"USDBRL":{"bid":"5.30"}}`
	if string(data) != expected {
		t.Errorf("Expected %v but got %v", expected, string(data))
	}

}

// unhappy path
func TestShouldRespondErrorWhenExchangeApiTimeoutHappen(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", ExchangeRateApiURL,
		httpmock.NewStringResponder(200, "{}").Delay(201*time.Millisecond))

	expected := `Get "https://economia.awesomeapi.com.br/json/last/USD-BRL": context deadline exceeded`

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	CotationHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code 500 but got %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected %v but got %v", expected, string(data))
	}

}

func TestShouldRespondErrorWhenExchangeApiFails(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", ExchangeRateApiURL,
		httpmock.NewErrorResponder(errors.New("Something went wrong")))

	expected := `Get "https://economia.awesomeapi.com.br/json/last/USD-BRL": Something went wrong`

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	CotationHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code 500 but got %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected %v but got %v", expected, string(data))
	}

}
