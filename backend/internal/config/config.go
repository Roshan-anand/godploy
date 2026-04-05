package config

import (
	"os"
)

type Config struct {
	Port             string
	SessionDataName  string
	SessionTokenName string
	EchoCtxUserKey   string
	JwtSecret        string
	WebUrl           string
	DbDir            string
	AppEnv           string
}

func LoadConfig() (*Config, error) {
	appEnv := os.Getenv("APP_ENV")
	jwtSecrect := os.Getenv("JWT_SECRET")
	webUrl := os.Getenv("WEB_URL")

	// TODO: load from env variable
	return &Config{
		Port:             "8080",
		SessionDataName:  "godploy_session_data",
		SessionTokenName: "godploy_session_token",
		EchoCtxUserKey:   "user_email",
		JwtSecret:        jwtSecrect,
		WebUrl:           webUrl,
		DbDir:            "data",
		AppEnv:           appEnv,
	}, nil
}
