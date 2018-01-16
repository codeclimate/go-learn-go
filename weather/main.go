package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Forecast struct {
	City string `json:"city"`
	Zip  string `json:"zip"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in / handler", r)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	zip := r.URL.Query().Get("zip")

	forecast := Forecast{
		City: "Brooklyn",
		Zip:  zip,
	}

	b, err := json.Marshal(forecast)

	if err == nil {
		w.Write(b)
	} else {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", home)
	log.Print("Server started, yo!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
