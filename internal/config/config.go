package config

type Config struct {
	Port             string
	SessionDataName  string
	SessionTokenName string
	EchoCtxUserKey   string
	JwtSecret        string
	AllowedCors      []string
	DbDir            string
}

func LoadConfig() *Config {
	// TODO: load from env variable
	return &Config{
		Port:             "8080",
		SessionDataName:  "godploy_session_data",
		SessionTokenName: "godploy_session_token",
		EchoCtxUserKey:   "user_email",
		JwtSecret:        "my_secret",
		AllowedCors:      []string{"http://localhost:5173"},
		DbDir:            "data",
	}
}
