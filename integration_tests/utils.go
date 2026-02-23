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
	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/routes"
	"github.com/labstack/echo/v5"
)

// initialize a mock server for testing with config values suitable for testing
func mockConfigServer() (*echo.Echo, *config.Server, error) {

	cfg := config.LoadConfig()

	// update config to include testing data
	cfg.AllowedCors = []string{"*"}

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

	return r, server, nil
}

// mock db connection
func mockDbConnection() (*db.Queries, error) {
	tempDir, err := getTempDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get temp dir: %w", err)
	}

	db, err := config.InitDb(tempDir)
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}

	return db.Queries, nil
}

// mock a new logined user
func mockUserRejister(url string, h *http.Client, cfg *config.Config) (*routes.AuthRes, error) {
	rRegister := "/api/auth/register"

	registerReq := routes.RegisterReq{
		Name:     "test",
		Email:    "test@test.com",
		Password: "testtest",
		Org:      "test_org",
	}
	r, err := h.Post(url+rRegister, "application/json", reqBody(registerReq))
	if err != nil {
		return nil, fmt.Errorf("err making request: %v", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status code %d, got %d", http.StatusUnauthorized, r.StatusCode)
	}

	if !hasCookie(r.Cookies(), cfg) {
		return nil, fmt.Errorf("expected cookies not found in response")
	}

	body := new(routes.AuthRes)
	if err := readAndUnmarshl(r.Body, body); err != nil {
		return nil, fmt.Errorf("err reading response body: %v", err)
	}

	return body, nil
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

// reads the reader and unmarshal it
func readAndUnmarshl(body io.ReadCloser, v any) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	return nil
}

// check if cookies exists
func hasCookie(c []*http.Cookie, cfg *config.Config) bool {
	for _, cookie := range c {
		switch cookie.Name {
		case cfg.SessionDataName, cfg.SessionTokenName:
		default:
			return false
		}
	}
	return true
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
