package config

type Config struct {
	Port             string
	SessionDataName  string
	SessionTokenName string
	EchoCtxUserKey   string
	JwtSecret       string
}

func LoadConfig() *Config {
	return &Config{
		Port:             "8080",
		SessionDataName:  "godploy_session_data",
		SessionTokenName: "godploy_session_token",
		EchoCtxUserKey:   "user_email",
		JwtSecret: "my_secret", // TODO: load from env variable
	}
}
