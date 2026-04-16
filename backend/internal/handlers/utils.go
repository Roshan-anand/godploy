package handlers

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

// Generates a unique GitHub App name per manifest to avoid collisions across setup attempts.
func generateGitHubManifestAppName() (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	randomBytes := make([]byte, 6)
	suffix := make([]byte, 6)

	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	for i, b := range randomBytes {
		suffix[i] = chars[int(b)%len(chars)]
	}

	return "godploy-" + string(suffix), nil
}

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
func getManifestData(url string, state string) (string, error) {
	appName, err := generateGitHubManifestAppName()
	if err != nil {
		return "", err
	}

	manifest := map[string]interface{}{
		"name": appName,
		"url":  url,
		"hook_attributes": map[string]string{
			"url": "https://example.com/github/events", // TODO : replace with webhook endpoint URL
		},
		"redirect_url": url + "/api/provider/github/app/callback",
		// "callback_urls": []string{"http://localhost:8080/api/provider/github/app/callback"},
		"setup_url": fmt.Sprintf("%s/api/provider/github/app/setup?state=%s", url, state),
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
