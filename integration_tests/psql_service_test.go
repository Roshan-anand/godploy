package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/routes"
)

func TestPsqlOperation(t *testing.T) {

	// initialize mock server
	e, srv, err := mockConfigServer()
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

	user, err := mockUserRejister(ts.URL, h, srv.Config)
	if err != nil {
		t.Fatal(err)
	}
	orgid := user.Orgs[0].ID

	// route req body
	p := new(db.CreateProjectRow)
	r, err := h.Post(ts.URL+"/api/project", "application/json", reqBody(routes.CreateProjectReq{Name: "test", OrgID: orgid}))
	if err != nil {
		t.Fatal("err making request:", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, r.StatusCode)
	}

	if err := readAndUnmarshl(r.Body, p); err != nil {
		t.Fatal(err)
	}

	// route paths
	rPsql := "/api/service/psql"

	// route req body
	rCreateBody := routes.CreatePsqlServiceReq{ProjectID: p.ID, Name: "test", AppName: "test", Description: "", DbName: "test", DbUser: "test", DbPassword: "test", Image: "postgres:16"}
	rPsqlBody := routes.PsqlServiceReq{PsqlId: "test"}

	// route res body
	createBodyRes := new(db.PsqlService)

	t.Run("POST /service/psql : returns sucess for creating service", func(t *testing.T) {
		r, err := h.Post(ts.URL+rPsql, "application/json", reqBody(rCreateBody))
		if err != nil {
			t.Fatal("err making request:", err)
		}
		defer r.Body.Close()

		if r.StatusCode != http.StatusOK {
			t.Fatalf("expected status code %d, got %d", http.StatusOK, r.StatusCode)
		}

		if err := readAndUnmarshl(r.Body, createBodyRes); err != nil {
			t.Fatal(err)
		}
	})

	rPsqlBody.PsqlId = createBodyRes.PsqlID

	t.Run("POST /service/psql : returns sucess for creating service", func(t *testing.T) {
		deleteReq, err := getDeleteReq(ts.URL+rPsql, rPsqlBody)
		if err != nil {
			t.Fatal("err creating delete req:", err)
		}
		r, err := h.Do(deleteReq)
		if err != nil {
			t.Fatal("err making request:", err)
		}

		if r.StatusCode != http.StatusOK {
			t.Fatalf("expected status code %d, got %d", http.StatusOK, r.StatusCode)
		}
	})
}
