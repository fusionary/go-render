package render

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateResource(t *testing.T) {
	godotenv.Load(".env")

	token := os.Getenv("TOKEN")

	ctx := context.Background()
	ctx = StoreAuthentication(ctx, token)

	got := CreateResource(ctx)
	if got == false {
		t.Errorf("CreateResource %v != true", got)
	}
}
