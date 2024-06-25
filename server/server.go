package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type ExchangeRate struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func main() {
	db, err := sql.Open("sqlite3", "./quotes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS quotes (id INTEGER PRIMARY KEY, bid TEXT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatal(err)
	}

	// Define the HTTP handler for /cotacao endpoint
	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()

		resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
		if err != nil {
			http.Error(w, "Failed to get exchange rate", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var rate ExchangeRate
		if err := json.NewDecoder(resp.Body).Decode(&rate); err != nil {
			http.Error(w, "Failed to decode exchange rate", http.StatusInternalServerError)
			return
		}

		ctxDB, cancelDB := context.WithTimeout(ctx, 10*time.Millisecond)
		defer cancelDB()

		_, err = db.ExecContext(ctxDB, "INSERT INTO quotes (bid) VALUES (?)", rate.USDBRL.Bid)
		if err != nil {
			http.Error(w, "Failed to save quote", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(rate.USDBRL)
	})

	// Log that the service is starting
	log.Println("Starting server on :8080...")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
