package testing

import (
	"context"
	"testing"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/Roshan-anand/godploy/internal/types"
)

func TestDbQueries(t *testing.T) {
	q, err := mockDbConnection()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	email := "a@email.com"
	orgName := "org1"

	t.Run("create and get user", func(t *testing.T) {
		if _, err := q.CreateUser(ctx, db.CreateUserParams{
			Name:     "a",
			Email:    email,
			HashPass: "qwe",
			Role:     types.AdminRole,
		}); err != nil {
			t.Fatal("error create user :", err)
		}

		orgId, err := q.CreateOrg(ctx, db.CreateOrgParams{
			ID:   lib.NewID(),
			Name: orgName,
		})
		if err != nil {
			t.Fatal("error create org :", err)
		}

		if err := q.LinkUserNOrg(ctx, db.LinkUserNOrgParams{
			UserEmail:      email,
			OrganizationID: orgId,
		}); err != nil {
			t.Fatal("error link user n org :", err)
		}

		if _, err := q.GetUserByEmail(ctx, email); err != nil {
			t.Fatal("error get user :", err)
		}
	})
}
