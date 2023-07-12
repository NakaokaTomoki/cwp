package cwp

type Config struct {
	Token   string
	// RunMode Mode
}

// type Mode int

// const (
// 	GetWeather Mode = iota + 1
// 	// List
// 	// ListGroup
// 	// Delete
// 	// QRCode
// )

// func NewConfig(token string, mode Mode) *Config {
func NewConfig(token string) *Config {
	// return &Config{Token: token, RunMode: mode}
	return &Config{Token: token}
}

// func (m Mode) String() string {
func String() string {
	return "getweather"
}
