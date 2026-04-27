package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port             string
	SessionDataName  string
	SessionTokenName string
	EchoCtxUserKey   string
	JwtSecret        string
	WebUrl           string
	ServerUrl        string
	SqliteDir        string
	BadgerDir        string
	AppEnv           string
}

func LoadConfig() (*Config, error) {
	appEnv := os.Getenv("APP_ENV")
	jwtSecrect := os.Getenv("JWT_SECRET")
	webUrl := os.Getenv("WEB_URL")
	srvUrl := os.Getenv("SERVER_PUBLIC_URL")

	fmt.Println("server url : ", srvUrl)

	// TODO : load from env variable
	return &Config{
		Port:             "8080",
		SessionDataName:  "godploy_session_data",
		SessionTokenName: "godploy_session_token",
		EchoCtxUserKey:   "user_email",
		JwtSecret:        jwtSecrect,
		WebUrl:           webUrl,
		SqliteDir:        "data/sqlite",
		BadgerDir:        "data/badger",
		AppEnv:           appEnv,
		ServerUrl:        srvUrl,
	}, nil
}
