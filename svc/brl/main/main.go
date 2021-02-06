package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://api.exchangeratesapi.io/latest?symbols=BRL")
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		content := string(bytes)
		log.Print(content)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, content)
	})

	http.ListenAndServe("0.0.0.0:8080", nil)
}
