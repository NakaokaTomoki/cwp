package cwp

import "fmt"

type Weather struct {
 	Place        string
	// Shorten   string `json:"link"`
	// Original  string `json:"long_url"`
	// IsDeleted bool   `json:"is_deleted"`
	// Group     string
}

func (weather *Weather) String() string {
	return fmt.Sprintf("%s", weather.Place)
}

type CurrentWeatherPrinter interface {
	// List(config *Config) ([]*Weather, error)
	// Shorten(config *Config, url string) (*Weather, error)
	Weather(config *Config, place string) (*Weather, error)
	// Delete(config *Config, Weather string) error
	// QRCode(config *Config, Weather string) ([]byte, error)
}
