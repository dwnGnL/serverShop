package apiserver

// Config ...
type Config struct {
	BindAdress string `toml:"bind_addr"`
	LogLevel   string `toml:"log_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAdress: ":8080",
		LogLevel:   "debug",
	}
}
