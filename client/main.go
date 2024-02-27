package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-fullcicle/client/config"
)

func main() {
	config.ParseConfig()
	bid, err := getDollarBid()
	if err != nil {
		log.Fatalf("Failed to get dollar bid from server, %v", err)
	}

	err = saveBid(bid)
	if err != nil {
		log.Fatalf("Failed to save dollar bid on file, %v", err)
	}
}

// getDollarBid uses the app's server to get the current dollar bid
func getDollarBid() (bid string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	u := fmt.Sprintf("%s/cotacao", config.GetConfig().ServerURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		if urlErr, ok := err.(*url.Error); ok && urlErr.Timeout() {
			err = errors.New("Request for server timed out after 200 milliseconds")
		}
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Failed to get dollar bid with status code: %d", res.StatusCode)
		return
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&bid)
	if err != nil {
		return
	}

	return
}

// saveBid saves the dollar bid to a file
func saveBid(bid string) (err error) {
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	content := fmt.Sprintf("Dólar:%s\n", bid)
	if _, err = file.WriteString(content); err != nil {
		return
	}

	return
}
