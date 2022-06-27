package render

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func CreateJob(ctx context.Context, serviceId string) bool {
	url := fmt.Sprintf("https://api.render.com/v1/services/%v/jobs", serviceId)

	payload := strings.NewReader("{\"startCommand\":\"node index.js\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", ctx.Value("token")))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 200 || res.StatusCode == 201 {
		return true
	} else {
		return false
	}
}
