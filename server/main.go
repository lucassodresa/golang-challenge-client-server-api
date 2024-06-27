package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := ConnectToDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = InitMigrations(db)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/cotation", CotationHandler)
	http.ListenAndServe(":8080", mux)
}

func LogError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
	fmt.Println(err)
}

func CotationHandler(w http.ResponseWriter, r *http.Request) {
	data, err := GetExchangeRate()
	if err != nil {
		LogError(w, err)
		return
	}

	db, err := ConnectToDB()
	if err != nil {
		LogError(w, err)
		return
	}
	defer db.Close()

	err = SaveExchangeRate(db, data.USDBRL.Bid)
	if err != nil {
		LogError(w, err)
		return
	}

	response, err := json.Marshal(data.USDBRL)
	if err != nil {
		LogError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
