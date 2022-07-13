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

	client := NewClient(nil, token)
	options := new(ServiceListOptions)
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
	client := NewClient(nil, token)
	options := new(ResourceGetEnvOptions)
	data, res, err := client.Services.ServiceGetEnvVars(ctx, serviceId, options)
	if err != nil {
		fmt.Printf("Error while getting results")
		fmt.Println(err)
	} else {
		log.Printf("%d %d", res.StatusCode, len(*data))
	}
}

func TestCreateService(t *testing.T) {
	godotenv.Load(".env")
	token := os.Getenv("TOKEN")

	ctx := context.Background()

	body := ServiceCreateBody{AutoDeploy: No, ServiceType: WebService, Name: "test-web-service", OwnerId: "tea-c8u8ehfd17c7d61uhjk0", Repo: "https://github.com/fusionary/digitalrealty.com.git", ServiceDetails: ServiceDetails{PullRequestPreviewsEnabled: No, Env: Docker}}

	client := NewClient(nil, token)
	data, res, err := client.Services.CreateService(ctx, &body)

	if err != nil {
		fmt.Printf("Error while getting results")
		fmt.Println(err)
	} else {
		log.Printf("%d %s", res.StatusCode, *data)
	}
}
