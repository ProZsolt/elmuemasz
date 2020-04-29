package elmuemasz

import (
	"os"
	"testing"
	"time"
)

func TestThings(t *testing.T) {
	username := os.Getenv("ELMU_USERNAME")
	password := os.Getenv("ELMU_PASSWORD")

	user := User{Username: username,
		Password: password,
	}

	ac, err := getAuthCode(user)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	token, err := getToken(ac)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	time.Sleep(time.Duration(10) * time.Second)

	token, err = renewToken(token)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// ctx := context.Background()
	// ts := oauth2.StaticTokenSource(&token)
	// client := oauth2.NewClient(ctx, ts)
	// _, err client.Get("http://example.com")
	// if err != nil {
	// 	t.Fatalf("Unexpected error: %v", err)
	// }
}
