package cwp

import (
	"encoding/json"
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
	"net/url"
	// "strings"
        // "time"

	// "os"
)

type OpenWeatherMap struct {
	url   string
}

func NewOpenWeatherMap() *OpenWeatherMap {
	return &OpenWeatherMap{url: "https://api.openweathermap.org/data/2.5/weather"}
}

// func (openWeathermap *OpenWeatherMap) List(config *Config) ([]*OpenWeather, error) {
// 	data, err := sendRequest(openWeathermap, place, config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return handleListResponse(data)
// }
//
// func handleListResponse(data []byte) ([]*OpenWeather, error) {
// 	result := struct {
// 		Links []*OpenWeather `json:"links"`
// 	}{}
// 	err := json.Unmarshal(data, &result)
// 	return nil, err
// }

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
	// fmt.Println()

	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	// fmt.Println(result.Place)
	// fmt.Println(result.Weather[0].WeatherDescription)
	// fmt.Println(result.Temp["temp"])
	// fmt.Println(time.Unix(result.DateTime, 0))

	return result, err
}

// func (openWeathermap *OpenWeatherMap) Delete(config *Config, shortenURL string) error {
// 	request, err := http.NewRequest("DELETE", openWeathermap.buildUrl("bitlinks/"+strings.TrimPrefix(shortenURL, "https://")), nil)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = sendRequest(request, config)
// 	return err
// }

// func (openWeathermap *OpenWeatherMap) QRCode(config *Config, shortenURL string) ([]byte, error) {
// 	return []byte{}, fmt.Errorf("not implement yet")
// }

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

	// return handleResponse(response)
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
