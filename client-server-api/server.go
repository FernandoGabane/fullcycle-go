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

type CotacaoResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS cotacao (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		ctxAPI, cancelAPI := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancelAPI()

		req, err := http.NewRequestWithContext(ctxAPI, http.MethodGet,
			"https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("API timeout or error:", err)
			w.WriteHeader(http.StatusGatewayTimeout)
			return
		}
		defer resp.Body.Close()

		var cotacao CotacaoResponse
		if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancelDB()

		_, err = db.ExecContext(ctxDB,
			"INSERT INTO cotacao (bid) VALUES (?)",
			cotacao.USDBRL.Bid,
		)
		if err != nil {
			log.Println("DB timeout or error:", err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"bid": cotacao.USDBRL.Bid,
		})
	})

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
