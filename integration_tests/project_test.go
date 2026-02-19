package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/routes"
)

func TestProjectOperations(t *testing.T) {
	// route paths
	rCreateProject := "/api/project"

	createBody := routes.CreateProjectReq{Name: "new"}

	// initialize mock server
	e, cfg, err := mockConfigServer()
	if err != nil {
		t.Fatal("err config server :", err)
	}

	// start test server
	ts := httptest.NewServer(e)
	// url, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal("err parsing url:", err)
	}
	t.Cleanup(ts.Close)

	// create  new global client
	h, err := getNewClient()
	if err != nil {
		t.Fatal("err creating http client:", err)
	}

	if err := mockUserRejister(ts.URL, h, cfg); err != nil {
		t.Fatal(err)
	}

	var body db.CreateProjectRow

	t.Run("/project : returns sucess for creating project", func(t *testing.T) {
		r, err := h.Post(ts.URL+rCreateProject, "application/json", reqBody(createBody))
		if err != nil {
			t.Fatal("err making request:", err)
		}
		defer r.Body.Close()

		if err := readAndUnmarshl(r.Body, &body); err != nil {
			t.Fatal(err)
		}

		if r.StatusCode != http.StatusOK {
			t.Fatalf("expected status code %d, got %d", http.StatusOK, r.StatusCode)
		}
	})
}
