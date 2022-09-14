package render

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type JobsService service

type Job struct {
	Id           string `json:"id,omitempty"`
	ServiceId    string `json:"serviceId,omitempty"`
	StartCommand string `json:"startCommand,omitempty"`
	PlanId       string `json:"planId,omitempty"`
	Status       string `json:"status,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
	StartedAt    string `json:"startedAt,omitempty"`
	FinishedAt   string `json:"finishedAt,omitempty"`
}

type JobCreateBody struct {
	StartCommand string `json:"startCommand,omitempty"`
	PlanId       string `json:"planId,omitempty"`
}

func (s *JobsService) CreateJob(ctx context.Context, serviceId string, jobCreateBody JobCreateBody) (*Job, *http.Response, error) {
	url := fmt.Sprintf("services/%s/jobs", serviceId)
	log.Println(url)
	req, err := s.client.NewRequest("POST", url, jobCreateBody)
	if err != nil {
		return nil, nil, err
	}

	var job *Job
	res, err := s.client.Do(ctx, req, &job)
	if err != nil {
		return nil, nil, err
	}

	return job, res, nil
}

func (s *DeploysService) RetrieveJob(ctx context.Context, serviceId string, jobId string) (*Job, *http.Response, error) {
	url := fmt.Sprintf("services/%s/jobs/%s", serviceId, jobId)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var job *Job
	res, err := s.client.Do(ctx, req, &job)
	if err != nil {
		return nil, nil, err
	}

	return job, res, nil
}
