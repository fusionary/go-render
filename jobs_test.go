package render

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateJob(t *testing.T) {
	godotenv.Load(".env")

	token := os.Getenv("TOKEN")

	ctx := context.Background()
	ctx = StoreAuthentication(ctx, token)

	got := CreateJob(ctx, os.Getenv("JOB_TEST_SERVICEID"))
	if got == false {
		t.Errorf("CreateJob %v != true", got)
	}
}
