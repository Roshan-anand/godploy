package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

// remove the session data
func removeSession(query *db.Queries, state string) {
	if err := query.DeleteRedirectSession(context.Background(), state); err != nil {
		fmt.Println("Error deleting redirect session:", err)
	}
}

// binds and validate the given data
func BindAndValidate(b any, c *echo.Context, v *validator.Validate) *lib.Res {

	if err := c.Bind(b); err != nil {
		return &lib.Res{Message: "Invalid Data"}
	}

	if err := v.Struct(b); err != nil {
		return &lib.Res{Message: fmt.Sprintf("validation error : %v", err)}
	}

	return nil
}

// get github app manifest data
func getManifestData(url string, orgId uuid.UUID) (string, error) {
	manifest := map[string]interface{}{
		"name": "GODPLOY",
		"url":  url,
		"hook_attributes": map[string]string{
			"url": "https://example.com/github/events", // TODO : replace with webhook endpoint URL
		},
		"redirect_url": url + "/api/provider/github/app/callback",
		// "callback_urls": []string{"http://localhost:8080/api/provider/github/app/callback"},
		"setup_url": fmt.Sprintf("%s/api/provider/github/app/setup?org_id=%s", url, orgId.String()),
		"public":    true,
		"default_permissions": map[string]string{
			"contents": "read",
			"metadata": "read",
		},
		"default_events": []string{"push"},
	}

	manifestDataB, err := json.Marshal(manifest)
	if err != nil {
		return "", err
	}

	return string(manifestDataB), nil
}

// auto-submitting form template — POST to GitHub with manifest in body (required by GitHub manifest flow)
const githubManifestFormTmpl = `<!DOCTYPE html>
<html>
<body>
  <form id="mf" action="https://github.com/settings/apps/new?state={{.State}}" method="POST">
    <input type="hidden" name="manifest" value="{{.Manifest}}">
  </form>
  <script>document.getElementById("mf").submit();</script>
</body>
</html>`
