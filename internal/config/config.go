package config

type Config struct {
	Port             string
	SessionDataName  string
	SessionTokenName string
}

func LoadConfig() *Config {
	return &Config{
		Port:             "8080",
		SessionDataName:  "godploy_session_data",
		SessionTokenName: "godploy_session_token",
	}
}
