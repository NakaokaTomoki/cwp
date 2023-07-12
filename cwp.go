package cwp

import "fmt"

type OpenWeather struct {
	Place              string    `json:"name"`
	Weather            []Weather `json:"weather"`
        Temp               map[string]interface{}   `json:"main"`
	DateTime           int64     `json:"dt"`
	// Shorten   string `json:"link"`
	// Original  string `json:"long_url"`
	// IsDeleted bool   `json:"is_deleted"`
	// Group     string
}

type Weather struct {
	Weather            string    `json:"main"`
	WeatherDescription string    `json:"description"`
}

type Main struct {
	Main            float64    `json:"main"`
}

func (openweather *OpenWeather) String() string {
	return fmt.Sprintf("%s", openweather.Place)
}

type CurrentWeatherPrinter interface {
	// List(config *Config) ([]*Weather, error)
	// Shorten(config *Config, url string) (*Weather, error)
	OpenWeather(config *Config, place string) (*OpenWeather, error)
	// Delete(config *Config, Weather string) error
	// QRCode(config *Config, Weather string) ([]byte, error)
}
