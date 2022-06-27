package render

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func CreateResource(ctx context.Context) bool {
	url := "https://api.render.com/v1/services"

	payload := strings.NewReader("{\"autoDeploy\":\"no\",\"serviceDetails\":{\"pullRequestPreviewsEnabled\":\"no\",\"disk\":{\"sizeGB\":10,\"name\":\"digitalrealty-database-disk-prod\",\"mountPath\":\"/var/lib/mysql\"},\"envSpecificDetails\":{\"dockerfilePath\":\"./docker/Dockerfile.mysql\",\"dockerContext\":\".\"},\"numInstances\":1,\"plan\":\"starter_plus\",\"region\":\"oregon\",\"env\":\"docker\"},\"type\":\"private_service\",\"repo\":\"https://github.com/fusionary/digitalrealty.com.git\",\"name\":\" digitalrealty-database-prod\",\"ownerId\":\"tea-c8u8ehfd17c7d61uhjk0\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", ctx.Value("token")))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 201 {
		return true
	} else {
		log.Printf("%d %s", res.StatusCode, body)
		return false
	}
}
