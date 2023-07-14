package cwp

import (
	// "os"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	// "strings"
        // "time"
)

type OpenWeatherMap struct {
	url   string
}

func NewOpenWeatherMap() *OpenWeatherMap {
	return &OpenWeatherMap{url: "https://api.openweathermap.org/data/2.5/weather"}
}

func (openWeathermap *OpenWeatherMap) GetWeather(place string, config *Config) (*OpenWeather, error) {
	data, err := sendRequest(openWeathermap, place, config)
	if err != nil {
		return nil, err
	}
	return handleApiResponse(data)
}

func handleApiResponse(data []byte) (*OpenWeather, error) {
	result := &OpenWeather{}
	// fmt.Println("result:", string(data))

	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func sendRequest(openWeathermap *OpenWeatherMap, place string, config *Config) ([]byte, error) {
	params := url.Values{}
	params.Set("q", place)
	params.Set("appid", config.Token)
	params.Set("units", "metric")
	params.Set("lang", "ja")

	response, err := http.Get(openWeathermap.url + "?" + params.Encode())
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

        bytes, err := handleResponse(response)
	if err != nil {
		return []byte{}, err
	}
	return bytes, err
}

func handleResponse(response *http.Response) ([]byte, error) {
	if response.StatusCode/100 != 2 {
		data, _ := io.ReadAll(response.Body)
		fmt.Println("response body:", string(data))
		return []byte{}, fmt.Errorf("response status code %d", response.StatusCode)
	}
	return io.ReadAll(response.Body)
}
