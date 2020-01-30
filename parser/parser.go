package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/

const apiEndpoint = "https://api.openweathermap.org/data/2.5/weather?zip="

func parseJSON(body []byte) OpenWeatherResponse {
	var openWeatherResponse OpenWeatherResponse
	json.Unmarshal(body, &openWeatherResponse)
	return openWeatherResponse
}

func errorCheck(e error, msg string) {
	if e != nil {
		fmt.Println(msg)
		fmt.Println(e.Error())
	}
}

// Runme is a Dank function, I can't believe my linter wants comments
func Runme() int {
	apiKey := os.Getenv("WEATHER_KEY")

	if apiKey == "" {
		fmt.Println("no key found")
		return -1
	}

	fmt.Println(apiKey)
	zipcode := 98105
	url := fmt.Sprintf("%s%d,us&APPID=%s", apiEndpoint, zipcode, apiKey)
	fmt.Println(url)
	resp, err := http.Get(url)

	errorCheck(err, "error in get request")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	errorCheck(err, "error in reading request body")

	fmt.Println(string(body))
	openWeatherResponse := parseJSON(body)

	/* Example using interfaces and methods */
	// Turns out that the OpenWeatherResponse has the method attached to it?
	fmt.Println(openWeatherResponse.GetJSONResponse())

	return 0
}
