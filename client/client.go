package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dmarins/client-server-challenge-go/domain"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var price domain.Price

	err = json.Unmarshal(body, &price)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	if string(price.Bid) != "" {
		file, err := os.Create("./client/cotacao.txt")
		if err != nil {
			log.Println(err)
			panic(err)
		}

		defer file.Close()

		message := fmt.Sprintf("Dólar: %s", price.Bid)

		_, err = file.Write([]byte(message))
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}
