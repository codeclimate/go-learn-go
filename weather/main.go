package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Forecast struct {
	City        string `json:"city"`
	Zip         string `json:"zip"`
	Description string `json:"description"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json"description"`
}

type OpenWeatherMapResponse struct {
	Name    string    `json:"name"`
	Weather []Weather `json:"weather"`
}

func getWeather(zip string, api_key string) (weather OpenWeatherMapResponse, weather_error error) {
	weather_url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?zip=%s,us&appid=%s",
		zip,
		api_key,
	)

	resp, weather_error := http.Get(weather_url)
	if weather_error != nil {
		return
	}

	// TODO: new error here (maybe, later error catching will prob catch it)
	// if resp.StatusCode != 200 {
	// 	weather_error = err
	// 	return
	// }

	defer resp.Body.Close()
	bodyBytes, weather_error := ioutil.ReadAll(resp.Body)
	if weather_error != nil {
		return
	}

	weather = OpenWeatherMapResponse{}
	weather_error = json.Unmarshal(bodyBytes, &weather)

	return
}

func home(w http.ResponseWriter, r *http.Request) {
	// TODO: load from .env file or something
	open_weather_map_api_key := os.Getenv("OPEN_WEATHER_MAP_API_KEY")

	fmt.Println("in / handler", r)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	zip := r.URL.Query().Get("zip")

	weather, err := getWeather(zip, open_weather_map_api_key)

	if err != nil {
		log.Fatal(err)
	}

	forecast := Forecast{
		City:        weather.Name,
		Zip:         zip,
		Description: weather.Weather[0].Description,
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
