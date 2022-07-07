package render

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type DeploysService service

type Commit struct {
	Id        *string    `json:"id,omitempty"`
	Message   *string    `json:"message,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

type Deploy struct {
	Id         *string    `json:"id,omitempty"`
	Commit     *Commit    `json:"commit,omitempty"`
	Status     *string    `json:"status,omitempty"`
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	CpdatedAt  *time.Time `json:"updatedAt,omitempty"`
}

type DeployTriggerCacheClear string

const (
	DoNotClear DeployTriggerCacheClear = "do_not_clear"
	Clear      DeployTriggerCacheClear = "clear"
)

type DeployTriggerBody struct {
	ClearCache DeployTriggerCacheClear `json:"clearCache,omitempty"`
}

func (s *DeploysService) TriggerADeployment(ctx context.Context, serviceId string, deployTriggerBody DeployTriggerBody) (*Deploy, *http.Response, error) {
	url := fmt.Sprintf("services/%s/deploys", serviceId)

	req, err := s.client.NewRequest("POST", url, deployTriggerBody)
	if err != nil {
		return nil, nil, err
	}

	var deploy *Deploy
	res, err := s.client.Do(ctx, req, &deploy)
	if err != nil {
		return nil, nil, err
	}

	return deploy, res, nil
}
