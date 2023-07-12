package cwp

import "fmt"

type OpenWeather struct {
	Place              string                   `json:"name"`
	Weather            []Weather                `json:"weather"`
        Temp               map[string]interface{}   `json:"main"`
	DateTime           int64                    `json:"dt"`
}

type Weather struct {
	Weather            string                   `json:"main"`
	WeatherDescription string                   `json:"description"`
}

type Main struct {
	Main               float64                  `json:"main"`
}

func (openweather *OpenWeather) String() string {
	return fmt.Sprintf("%s", openweather.Place)
}

type CurrentWeatherPrinter interface {
 	GetWeather (place string, config *Config) (*OpenWeather, error)
}
