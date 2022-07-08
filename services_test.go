package render

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestServiceList(t *testing.T) {
	godotenv.Load(".env")
	token := os.Getenv("TOKEN")

	ctx := context.Background()

	options := new(ServiceListOptions)
	client := NewClient(nil, token)
	data, res, err := client.Services.ListServices(ctx, options)
	if err != nil {
		fmt.Printf("Error while getting results")
		fmt.Println(err)
	} else {
		log.Printf("%d %d", res.StatusCode, len(*data))
	}
}

func TestServiceGetEnvVars(t *testing.T) {
	godotenv.Load(".env")
	token := os.Getenv("TOKEN")

	ctx := context.Background()

	serviceId := os.Getenv("ENV_TEST_SERVICEID")
	options := new(ResourceGetEnvOptions)
	client := NewClient(nil, token)
	data, res, err := client.Services.ServiceGetEnvVars(ctx, serviceId, options)
	if err != nil {
		fmt.Printf("Error while getting results")
		fmt.Println(err)
	} else {
		log.Printf("%d %d", res.StatusCode, len(*data))
	}
}
