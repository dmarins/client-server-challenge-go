package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dmarins/client-server-challenge-go/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	http.HandleFunc("/cotacao", handler)
	log.Printf("Starting server at port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	price, err := getPrice(ctx)
	if err != nil {
		panic(err)
	}

	if price != nil && string(price.Bid) != "" {
		err = savePrice(ctx, price)
		if err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(price)
}

func getPrice(ctx context.Context) (*domain.Price, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
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

	var rawJson map[string]interface{}
	json.Unmarshal(body, &rawJson)

	strJson, _ := json.Marshal(rawJson["USDBRL"])

	var price domain.Price
	err = json.Unmarshal(strJson, &price)
	if err != nil {
		return nil, err
	}

	return &price, nil
}

func savePrice(ctx context.Context, price *domain.Price) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	db, err := initDb()
	if err != nil {
		return err
	}

	result := db.WithContext(ctx).Create(&price)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func initDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("price.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.Price{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
