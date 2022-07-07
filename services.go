package render

import (
	"context"
	"net/http"
	"time"
)

type ServicesService service

type ServiceType string

const (
	StaticSite       string = "static_site"
	WebService       string = "web_service"
	PrivateService   string = "private_service"
	BackgroundWorker string = "background_worker"
	CronJob          string = "cron_job"
)

type ServiceListOptions struct {
	Name          []string      `url:"name"`
	ServiceType   []ServiceType `url:"type"`
	Env           []string      `url:"env"`
	Region        []string      `url:"region"`
	Suspended     []string      `url:"suspended"`
	CreatedBefore time.Time     `url:"createdBefore"`
	CreatedAfter  time.Time     `url:"createdAfter"`
	UpdatedBefore time.Time     `url:"updatedBefore"`
	UpdatedAfter  time.Time     `url:"updatedAfter"`
	OwnerId       []string      `url:"ownerId"`
	Cursor        string        `url:"cursor"`
	Limit         int           `url:"limit"`
}

type Service struct {
	Id             string    `json:"id,omitempty"`
	AutoDeploy     string    `json:"autoDeploy,omitempty"`
	Branch         string    `json:"branch:omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
	Name           string    `json:"name,omitempty"`
	NotifyOnFail   string    `json:"notifyOnFail,omitempty"`
	OwnerId        string    `json:"ownerId,omitempty"`
	Repo           string    `json:"repo,omitempty"`
	Slug           string    `json:"slug,omitempty"`
	Suspended      string    `json:"suspended,omitempty"`
	Suspenders     []string  `json:"suspenders,omitempty"`
	ServiceType    string    `json:"type,omitempty"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty"`
	ServiceDetails any       `json:"serviceDetails,omitempty"`
}

type ServiceResponse struct {
	Service []*Service `json:"service,omitempty"`
	Cursor  string     `json:"string,omitempty"`
}

func (s *ServicesService) ListServices(ctx context.Context, opts *ServiceListOptions) (*[]ServiceResponse, *http.Response, error) {
	url := "services"
	url, err := addOptions(url, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var serviceResponse *[]ServiceResponse
	res, err := s.client.Do(ctx, req, &serviceResponse)
	if err != nil {
		return nil, nil, err
	}

	return serviceResponse, res, nil
}
