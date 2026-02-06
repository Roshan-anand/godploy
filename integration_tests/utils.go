package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"os"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/routes"
	"github.com/labstack/echo/v5"
)

// initialize a mock server for testing with config values suitable for testing
func mockConfigServer() (*echo.Echo, *config.Config, error) {

	cfg := config.LoadConfig()

	// update config to include testing data
	cfg.AllowedCors = []string{"*"}

	// port, err := GetFreePort()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get free port: %w", err)
	// }
	// cfg.Port = fmt.Sprintf(":%d", port)

	tempDir, err := getTempDir()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get temp dir: %w", err)
	}
	cfg.DbDir = tempDir

	// create server instance
	server, err := config.NewServer(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize server: %w", err)
	}

	// setup all routes
	r, err := routes.SetupRoutes(server)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to setup routes: %w", err)
	}

	// server.SetupHttp(r) // setup http server with routes

	return r, cfg, nil
}

// get a new http client with cookie jar
func getNewClient() (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar: %w", err)
	}
	// create  new global client
	h := http.DefaultClient
	h.Jar = jar
	return h, nil
}

// get a random free port
func getFreePort() (int, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port, nil
}

// get temp dir for testing
func getTempDir() (string, error) {
	p, err := os.MkdirTemp("", "godploy_test_*")
	if err != nil {
		return "", err
	}
	return p, nil
}

func reqBody(data any) io.Reader {
	jsonData, _ := json.Marshal(data)
	return bytes.NewReader(jsonData)
}
