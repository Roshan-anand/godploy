package testing

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/routes"
)

func getDeleteReq(url string, rDeletBody routes.DeleteProjectReq) (*http.Request, error) {
	deleteReq, err := http.NewRequest(http.MethodDelete, url, reqBody(rDeletBody))
	if err != nil {
		return nil, fmt.Errorf("error creating delete req : %w", err)
	}
	deleteReq.Header.Set("Content-Type", "application/json")

	return deleteReq, nil
}

func TestProjectOperations(t *testing.T) {

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

	// route paths
	rProject := "/api/project"

	// route req body
	rCreateBody := db.CreateProjectParams{Name: "test", OrgID: orgid}
	rDeletBody := routes.DeleteProjectReq{}

	// route res body
	createBodyRes := new(db.CreateProjectRow)
	getBodyRes := new([]db.GetAllProjectsRow)

	t.Run("POST /project : returns sucess for creating project", func(t *testing.T) {
		r, err := h.Post(ts.URL+rProject, "application/json", reqBody(rCreateBody))
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

	t.Run("POST /project : returns conflict as tried to create same project in the org", func(t *testing.T) {
		r, err := h.Post(ts.URL+rProject, "application/json", reqBody(rCreateBody))
		if err != nil {
			t.Fatal("err making request:", err)
		}

		if r.StatusCode != http.StatusConflict {
			t.Fatalf("expected status code %d, got %d", http.StatusConflict, r.StatusCode)
		}
	})

	t.Run("GET /project?org_id : returns projects avalible in the org", func(t *testing.T) {
		r, err := h.Get(ts.URL + rProject + "?org_id=" + orgid)
		if err != nil {
			t.Fatal("err making request:", err)
		}
		defer r.Body.Close()

		if r.StatusCode != http.StatusOK {
			t.Fatalf("expected status code %d, got %d", http.StatusOK, r.StatusCode)
		}

		if err := readAndUnmarshl(r.Body, getBodyRes); err != nil {
			t.Fatal(err)
		}

		if len(*getBodyRes) == 0 {
			t.Fatalf("body is empty array : %v", getBodyRes)
		}
	})

	service, err := srv.DB.Queries.CreatePsqlService(context.Background(), db.CreatePsqlServiceParams{
		PsqlID:      "123",
		ProjectID:   createBodyRes.ID,
		Name:        "sample",
		AppName:     "sample",
		DbName:      "sample",
		Description: "",
		DbUser:      "",
		DbPassword:  "",
		Image:       "",
		InternalUrl: "",
	})
	if err != nil {
		t.Fatal(" error creatign service : ", err)
	}

	rDeletBody.ID = createBodyRes.ID

	t.Run("DELETE project: returns conflict as service exist", func(t *testing.T) {
		deleteReq, err := getDeleteReq(ts.URL+rProject, rDeletBody)
		if err != nil {
			t.Fatal(err)
		}
		r, err := h.Do(deleteReq)
		if err != nil {
			t.Fatal("err making request:", err)
		}

		if r.StatusCode != http.StatusConflict {
			t.Fatalf("expected status code %d, got %d", http.StatusConflict, r.StatusCode)
		}
	})

	if err := srv.DB.Queries.DeletePsqlService(context.Background(), service.PsqlID); err != nil {
		t.Fatal("error deleting service :", err)
	}

	t.Run("DELETE project: returns success for deleting project", func(t *testing.T) {
		deleteReq, err := getDeleteReq(ts.URL+rProject, rDeletBody)
		if err != nil {
			t.Fatal(err)
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
