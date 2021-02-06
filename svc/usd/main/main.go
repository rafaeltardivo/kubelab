package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type rate struct {
	USD float64 `json:"USD"`
}

type USD struct {
	Rates rate `json:"rates"`
}

func main() {
	var payload USD
	resp, err := http.Get("https://api.exchangeratesapi.io/latest?symbols=USD")
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		print(err)
	}

	fmt.Printf("%f", payload.Rates.USD)
}
