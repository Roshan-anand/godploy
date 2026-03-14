package lib

import (
	"context"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v84/github"
)

// CreateGithubClient creates an installation-scoped GitHub client.
// Used for repo operations (list repos, clone, etc.) scoped to a specific installation.
// Automatically handles JWT → installation token exchange and refresh.
func CreateGithubClient(ctx context.Context, appID int64, installationID int64, pemKey []byte) (*github.Client, error) {
	itr, err := ghinstallation.New(
		http.DefaultTransport,
		appID,
		installationID,
		pemKey,
	)
	if err != nil {
		return nil, err
	}

	client := github.NewClient(&http.Client{Transport: itr})
	return client, nil
}

// CreateAppClient creates an app-level GitHub client authenticated as the GitHub App itself (JWT).
// Required for app-level API calls like GetInstallation, ListInstallations — these endpoints
// only accept a JWT, not an installation access token.
func CreateAppClient(appID int64, pemKey []byte) (*github.Client, error) {
	itr, err := ghinstallation.NewAppsTransport(
		http.DefaultTransport,
		appID,
		pemKey,
	)
	if err != nil {
		return nil, err
	}

	client := github.NewClient(&http.Client{Transport: itr})
	return client, nil
}
