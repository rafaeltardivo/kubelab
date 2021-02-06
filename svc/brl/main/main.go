package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type rate struct {
	BRL float64 `json:"BRL"`
}

type BRL struct {
	Rates rate `json:"rates"`
}

func main() {
	var payload BRL
	resp, err := http.Get("https://api.exchangeratesapi.io/latest?symbols=BRL")
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		print(err)
	}

	fmt.Printf("%f", payload.Rates.BRL)
}
