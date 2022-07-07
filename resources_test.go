package render

import (
	"context"
	"fmt"
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

func TestGetEnvVars(t *testing.T) {
	godotenv.Load(".env")

	token := os.Getenv("TOKEN")
	serviceId := os.Getenv("ENV_TEST_SERVICEID")

	ctx := context.Background()
	ctx = StoreAuthentication(ctx, token)

	got, err := GetEnvVars(ctx, serviceId, nil)
	if err != nil {
		fmt.Printf("Error while getting results")
		fmt.Println(err)
	}
	if got == nil {
		t.Error("GetEnvVars nil")
	}
}
