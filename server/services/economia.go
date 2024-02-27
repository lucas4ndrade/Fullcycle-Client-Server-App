package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type response struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

// Uses the economia API to get the current dollar bid
func GetDollarBid() (str string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		if urlErr, ok := err.(*url.Error); ok && urlErr.Timeout() {
			err = errors.New("Request for economia.awesomeapi timed out after 200 milliseconds")
		}
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Failed to get dollar bid with status code: %d", res.StatusCode)
		return
	}

	var bodyRes response
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&bodyRes)
	if err != nil {
		return
	}

	str = bodyRes.USDBRL.Bid
	return
}
