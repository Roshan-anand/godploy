package testing

import (
	"testing"

	"github.com/Roshan-anand/godploy/internal/lib"
)

func TestGenerateAndVerifyJWT(t *testing.T) {
	u := lib.AuthUser{
		Email: "test@test.com",
		Name:  "Test User",
	}

	secret := "testsecretkey"

	jwtToken, err := lib.GenerateJWT(u, secret)
	if err != nil {
		t.Fatalf("Failed to generate JWT: %v", err)
	}

	claims, err := lib.VerifyJWT(jwtToken, secret)
	if err != nil {
		t.Fatalf("Failed to verify JWT: %v", err)
	}

	if claims.AuthUser != u {
		t.Errorf("Expected user %v \n got %v", u, claims.AuthUser)
	}
}
