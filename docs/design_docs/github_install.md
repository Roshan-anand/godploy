# GitHub App — Installation Flow (Step 4)

This picks up after the manifest flow is done and app credentials (`app_id`, `pem`, `webhook_secret`, `client_secret`) are stored in DB.

---

## What is an Installation?

The manifest flow **creates the app** — it doesn't give it access to any repo.
Installation is when a user/org explicitly grants the app access to their repos and an `installation_id` is produced.

You need `installation_id` to:
- Generate installation access tokens (via `ghinstallation`)
- Make API calls scoped to that org/user's repos

---

## The Flow

### Step 1 — Redirect user to install the app

Send the user to:

```
https://github.com/apps/{APP_SLUG}/installations/new
```

- `APP_SLUG` is the `slug` field from the manifest conversion response (e.g. `octoapp`)
- This is a plain GET redirect — just `window.location.href` or an `<a>` tag
- GitHub shows the user a permission grant screen

For org-scoped install:
```
https://github.com/apps/{APP_SLUG}/installations/new/permissions?target_id={ORG_ID}
```

---

### Step 2 — User approves, GitHub redirects to `setup_url`

After the user clicks Install, GitHub redirects to the `setup_url` you defined in the manifest:

```
https://your-app.com/api/provider/github/app/setup?installation_id=12345678
```

`setup_url` is a valid manifest parameter — it just wasn't in the example JSON in `github_provider.md`. You need to add it to your manifest payload:

```json
{
  "setup_url": "https://your-app.com/api/provider/github/app/setup"
}
```

GitHub appends `installation_id` as a query param when redirecting.

> **Do not trust this `installation_id` directly.** It can be spoofed (GitHub's own warning).
> Verify it by generating a user access token and checking which installations that user has.

---

### Step 3 — Verify + store the installation_id

To verify the `installation_id` is legit:

1. Use the app's JWT (signed with `pem`) to call:
   ```
   GET https://api.github.com/app/installations/{installation_id}
   ```
   Authorization: `Bearer <JWT>`

2. Response gives you the installation details including `account.login` (org or user name) and `account.id`.

3. Match `account.id` to the org in your DB, store `installation_id` linked to that org.

That's it — the installation is verified and stored.

---

### Step 4 — From here, `ghinstallation` handles the rest

With `app_id` + `pem` + `installation_id` stored, you use `ghinstallation` to:

```go
itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, appID, installationID, "path/to/pem")
client := github.NewClient(&http.Client{Transport: itr})
// client is now scoped to that installation — makes API calls with auto-refreshed tokens
```

JWT → installation access token → API calls — all handled by the package.

---

## Callback URL vs Setup URL (important distinction)

| Field in manifest | When GitHub hits it | Carries |
|---|---|---|
| `setup_url` | After user **installs** the app | `?installation_id=...` |
| `callback_urls` | After user **authorizes** the app (OAuth user flow) | `?code=...&state=...` |

For Godploy's use case (server-to-server, no user OAuth), **only `setup_url` matters**.
`callback_urls` is for user-facing OAuth — not needed here unless you want to act on behalf of a specific user.

---

## URL Summary

| Purpose | URL |
|---|---|
| Redirect user to install | `https://github.com/apps/{slug}/installations/new` |
| GitHub posts back to | your `setup_url` with `?installation_id=...` |
| Verify installation (JWT auth) | `GET https://api.github.com/app/installations/{installation_id}` |
| List all installations for app | `GET https://api.github.com/app/installations` |
| Get org-specific installation | `GET https://api.github.com/orgs/{org}/installation` |

---

## Relevant Docs

| Topic | URL |
|---|---|
| About setup_url | https://docs.github.com/en/apps/creating-github-apps/setting-up-a-github-app/about-the-setup-url |
| Installing a GitHub App (user side) | https://docs.github.com/en/apps/using-github-apps/installing-a-github-app-from-a-third-party |
| REST: Get an installation | https://docs.github.com/en/rest/apps/apps#get-an-installation-for-the-authenticated-app |
| REST: List installations | https://docs.github.com/en/rest/apps/apps#list-installations-for-the-authenticated-app |
| Authenticating as installation | https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-as-a-github-app-installation |
| ghinstallation pkg | https://github.com/bradleyfalzon/ghinstallation |
