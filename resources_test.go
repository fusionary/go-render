package render

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateResource(t *testing.T) {
	godotenv.Load(".env")

	token := os.Getenv("TOKEN")

	ctx := context.Background()
	ctx = StoreAuthentication(ctx, token)

	payload := strings.NewReader(os.Getenv("RESOURCES_TEST_PAYLOAD"))

	got := CreateResource(ctx, payload)
	if got == false {
		t.Errorf("CreateResource %v != true", got)
	}
}
