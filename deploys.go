package deploys

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Commit struct {
	id        *string    `json:"id,omitempty"`
	message   *string    `json:"message,omitempty"`
	createdAt *time.Time `json:"createdAt,omitempty"`
}

type Deploy struct {
	id         *string    `json:"id,omitempty"`
	commit     *Commit    `json:"commit,omitempty"`
	status     *string    `json:"status,omitempty"`
	finishedAt *time.Time `json:"finishedAt,omitempty"`
	createdAt  *time.Time `json:"createdAt,omitempty"`
	updatedAt  *time.Time `json:"updatedAt,omitempty"`
}

func TriggerDeploy(ctx context.Context, serviceId string) *bool {
	url := fmt.Sprintf("https://api.render.com/v1/services/%v/deploys", serviceId)

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", ctx.Value("token")))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 201 {
		return true
	} else {
		return false
	}
}
