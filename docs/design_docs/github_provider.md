# Git Provider - Github App

## Overview

Github App Manifest is used to create a Github App which does openrations on behalf of users.

- Fetch user repositories
- Pull repo code
- Webhook for repo events

---

## Github App Manifest Flow **(Oauth + Manifest Api)**

- pkg's :
  - "github.com/google/go-github/" : For all Github openrations API call.
  - "github.com/bradleyfalzon/ghinstallation/v2" : Token management, req caching, etc. for Github client

- Request for new app
  - **POST** req :
    - personal : https://github.com/settings/apps/new
    - organization : https://github.com/organizations/ORGANIZATION/settings/apps/new
  - /?manifest=manifest_json
  ```json
  {
    "name": "Octoapp",
    "url": "home page of the app",
    "hook_attributes": {
      "url": "https://app.com/github/events"
    },
    "redirect_url": "https://app.com/path/to/redirect",
		"setup_url": "http://app.com/path/to/setup?org_id=<xyz>",
    "public": true,
    "default_permissions": {
      "issues": "write",
      "checks": "write"
    },
    "default_events": ["issues", "issue_comment", "check_suite", "check_run"]
  }
  ```
  - &state=CSRF_token
    - why CSRF (Cross Site Request Forgery) : as this flow containes Back n forth, state changing request, to prevent access to unauthorized users.
- User prompts for permission.
- Get redirected to https://redirect_url?code=a180b1a3d263c81bc6441d7b990bae27d4c10679&state=abc123(if state provided) **(code=temp_code)**
- For handshake **POST** to https://api.github.com/app-manifests/{CODE}/conversions
  - Response :
  ```json
  {
    "id": 1,
    "slug": "octoapp",
    "node_id": "MDxOkludGVncmF0aW9uMQ==",
    "owner": {
      "login": "github",
      "id": 1,
      "node_id": "MDEyOk9yZ2FuaXphdGlvbjE=",
      "url": "https://api.github.com/orgs/github",
      "repos_url": "https://api.github.com/orgs/github/repos",
      "events_url": "https://api.github.com/orgs/github/events",
      "avatar_url": "https://github.com/images/error/octocat_happy.gif",
      "gravatar_id": "",
      "html_url": "https://github.com/octocat",
      "followers_url": "https://api.github.com/users/octocat/followers",
      "following_url": "https://api.github.com/users/octocat/following{/other_user}",
      "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
      "organizations_url": "https://api.github.com/users/octocat/orgs",
      "received_events_url": "https://api.github.com/users/octocat/received_events",
      "type": "User",
      "site_admin": true
    },
    "name": "Octocat App",
    "description": "",
    "external_url": "https://app.com",
    "html_url": "https://github.com/apps/octoapp",
    "created_at": "2017-07-08T16:18:44-04:00",
    "updated_at": "2017-07-08T16:18:44-04:00",
    "permissions": {
      "metadata": "read",
      "contents": "read",
      "issues": "write",
      "single_file": "write"
    },
    "events": ["push", "pull_request"],
    "client_id": "Iv1.8a61f9b3a7aba766",
    "client_secret": "1726be1638095a19edd134c77bde3aa2ece1e5d8",
    "webhook_secret": "e340154128314309424b7c8e90325147d99fdafa",
    "pem": "RSA PRIVATE KEY"
  }
  ```
- store pem, client_id, client_secret, webhook_secret in DB for future use.
- Redirect user to install the app
  ```
  https://github.com/apps/{APP_SLUG}/installations/new
  ```
- User approves, GitHub redirects to `setup_url`  
  ```
  https://app.com/path/to/setup?installation_id=12345678
  ```
- store installation_id in DB for future use.
- use pem, app_id & installation_id with ghinstallation to create github client to make API openrations.