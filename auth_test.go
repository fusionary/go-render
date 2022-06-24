package render

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestStoreAuthentication(t *testing.T) {
	godotenv.Load(".env")

	token := os.Getenv("TOKEN")

	ctx := context.Background()
	got := StoreAuthentication(ctx, token)

	if got.Value("token") != token {
		t.Errorf("StoreAuthentication %v != %v", got.Value("token"), token)
	}
}
