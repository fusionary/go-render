package render

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestTriggerDeploy(t *testing.T) {
	godotenv.Load(".env")
	token := os.Getenv("TOKEN")

	ctx := context.Background()

	body := DeployTriggerBody{ClearCache: DoNotClear}
	client := NewClient(nil, token)
	data, res, err := client.Deploys.TriggerADeployment(ctx, os.Getenv("DEPLOY_TEST_SERVICEID"), body)
	if err != nil {
		fmt.Printf("Error while getting results")
		fmt.Println(err)
	} else {
		json, err := json.Marshal(*data)
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Printf("%d %s", res.StatusCode, json)
	}
}
