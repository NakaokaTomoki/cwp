package cwp

type Config struct {
	Token   string
	RunMode Mode
}

type Mode int

const (
	GetWeather Mode = iota + 1
	List
	ListGroup
	// Delete
	// QRCode
)

func NewConfig(token string, mode Mode) *Config {
	return &Config{Token: token, RunMode: mode}
}

func (m Mode) String() string {
	switch m {
	case GetWeather:
		return "getweather"
	// case List:
	// 	return "list"
	// case ListGroup:
	// 	return "listgroup"
	// case Delete:
	// 	return "delete"
	// case QRCode:
	// 	return "qrcode"
	default:
		return "unknown"
	}
}
