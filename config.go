package cwp

type Config struct {
	Token   string
}

func NewConfig(token string) *Config {
	return &Config{Token: token}
}

func String() string {
	return "getweather"
}
