package render

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestTriggerDeploy(t *testing.T) {
	godotenv.Load(".env")

	token := os.Getenv("TOKEN")

	ctx := context.Background()
	ctx = StoreAuthentication(ctx, token)

	got := TriggerDeploy(ctx, os.Getenv("DEPLOY_TEST_SERVICEID"))
	if got == false {
		t.Errorf("TriggerDeploy %v != true", got)
	}
}