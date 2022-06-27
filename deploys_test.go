package render

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// var token = flag.String("key", "", "Render API Key")

func TestTriggerDeploy(t *testing.T) {
	godotenv.Load(".env")

	token := os.Getenv("TOKEN")

	ctx := context.Background()
	ctx = StoreAuthentication(ctx, token)

	got := TriggerDeploy(ctx, "srv-cafkfujru51gvuepntb0")
	if got == false {
		t.Errorf("TriggerDeploy %v != true", got)
	}
}
