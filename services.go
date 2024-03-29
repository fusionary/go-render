package render

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type ServicesService service

type ServiceType string

const (
	StaticSite       ServiceType = "static_site"
	WebService       ServiceType = "web_service"
	PrivateService   ServiceType = "private_service"
	BackgroundWorker ServiceType = "background_worker"
	CronJob          ServiceType = "cron_job"
)

type YesNo string

const (
	Yes YesNo = "yes"
	No  YesNo = "no"
)

type RouteType string

const (
	Redirect RouteType = "redirect"
	Rewrite  RouteType = "rewrite"
)

type ServiceEnvironmentType string

const (
	Docker ServiceEnvironmentType = "docker"
	Elixir ServiceEnvironmentType = "elixir"
	Go     ServiceEnvironmentType = "go"
	Node   ServiceEnvironmentType = "node"
	Python ServiceEnvironmentType = "python"
	Ruby   ServiceEnvironmentType = "ruby"
	Rust   ServiceEnvironmentType = "rust"
)

type ServiceRegion string

const (
	Oregon    ServiceRegion = "oregon"
	Frankfort ServiceRegion = "frankfort"
	Ohio      ServiceRegion = "ohio"
)

type ServicePlan string

const (
	Starter      ServicePlan = "starter"
	StarterPlus  ServicePlan = "starter_plus"
	Standard     ServicePlan = "standard"
	StandardPlus ServicePlan = "standard_plus"
	Pro          ServicePlan = "pro"
	ProPlus      ServicePlan = "pro_plus"
	ProMax       ServicePlan = "pro_max"
	ProUltra     ServicePlan = "pro_ultra"
)

type SecretFile struct {
	Name     string `json:"name,omitempty"`
	Contents string `json:"contents,omitempty"`
}

type Header struct {
	Path  string `json:"path,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Route struct {
	Type        RouteType `json:"type,omitempty"`
	Source      string    `json:"source,omitempty"`
	Destination string    `json:"destination,omitempty"`
}

type ServiceDisk struct {
	Name      string `json:"name,omitempty"`
	MountPath string `json:"mountPath,omitempty"`
	SizeGB    int    `json:"sizeGB,omitempty"`
}

type ServiceEnvironmentDockerDetailsPost struct {
	DockerCommand  string `json:"dockerCommand,omitempty"`
	DockerContext  string `json:"dockerContext,omitempty"`
	DockerfilePath string `json:"dockerfilePath,omitempty"`
}

type SpecificEnvironmentNativeDetailsPost struct {
	BuildCommand string `json:"buildCommand,omitempty"`
	StartCommand string `json:"startCommand,omitempty"`
}

type ServiceListOptions struct {
	Name          []string      `url:"name,omitempty"`
	ServiceType   []ServiceType `url:"type,omitempty"`
	Env           []string      `url:"env,omitempty"`
	Region        []string      `url:"region,omitempty"`
	Suspended     []string      `url:"suspended,omitempty"`
	CreatedBefore time.Time     `url:"createdBefore,omitempty"`
	CreatedAfter  time.Time     `url:"createdAfter,omitempty"`
	UpdatedBefore time.Time     `url:"updatedBefore,omitempty"`
	UpdatedAfter  time.Time     `url:"updatedAfter,omitempty"`
	OwnerId       []string      `url:"ownerId,omitempty"`
	Cursor        string        `url:"cursor,omitempty"`
	Limit         int           `url:"limit,omitempty"`
}

type Service struct {
	Id             string         `json:"id,omitempty"`
	AutoDeploy     string         `json:"autoDeploy,omitempty"`
	Branch         string         `json:"branch:omitempty"`
	CreatedAt      time.Time      `json:"createdAt,omitempty"`
	Name           string         `json:"name,omitempty"`
	NotifyOnFail   string         `json:"notifyOnFail,omitempty"`
	OwnerId        string         `json:"ownerId,omitempty"`
	Repo           string         `json:"repo,omitempty"`
	Slug           string         `json:"slug,omitempty"`
	Suspended      string         `json:"suspended,omitempty"`
	Suspenders     []string       `json:"suspenders,omitempty"`
	ServiceType    ServiceType    `json:"type,omitempty"`
	UpdatedAt      time.Time      `json:"updatedAt,omitempty"`
	ServiceDetails ServiceDetails `json:"serviceDetails,omitempty"`
}

type CreatedService struct {
	Service  Service `json:"service,omitempty"`
	DeployId string  `json:"deployId,omitempty"`
}

type ServiceResponse struct {
	Service Service `json:"service,omitempty"`
	Cursor  string  `json:"cursor,omitempty"`
}

type ServiceCreateBody struct {
	ServiceType          ServiceType    `json:"type,omitempty"`
	Name                 string         `json:"name,omitempty"`
	OwnerId              string         `json:"ownerId,omitempty"`
	Repo                 string         `json:"repo,omitempty"`
	AutoDeploy           YesNo          `json:"autoDeploy,omitempty"`
	Branch               string         `json:"branch,omitempty"`
	EnvironmentVariables []EnvVar       `json:"envVars,omitempty"`
	SecretFiles          []SecretFile   `json:"secretFiles,omitempty"`
	ServiceDetails       ServiceDetails `json:"serviceDetails,omitempty"`
}

type ServiceStaticSiteDetailsPost struct {
	BuildCommand               string   `json:"buildCommand,omitempty"`
	Headers                    []Header `json:"headers,omitempty"`
	PublishPath                string   `json:"publishPath,omitempty"`
	PullRequestPreviewsEnabled YesNo    `json:"pullRequestPreviewsEnabled,omitempty"`
	Routes                     []Route  `json:"routes,omitempty"`
}

type ServiceWebServiceDetailsPost struct {
	Disk                       *ServiceDisk           `json:"disk,omitempty"`
	Environment                ServiceEnvironmentType `json:"env,omitempty"`
	EnvironmentSpecificDetails any                    `json:"envSpecificDetails,omitempty"`
	HealthCheckPath            string                 `json:"healthCheckPath,omitempty"`
	NumberOfInstances          int                    `json:"numInstances,omitempty"`
	Plan                       ServicePlan            `json:"plan,omitempty"`
	PullRequestPreviewsEnabled YesNo                  `json:"pullRequestPreviewsEnabled,omitempty"`
	Region                     ServiceRegion          `json:"region,omitempty"`
}

type ServicePrivateServiceDetailsPost struct {
	Disk                       ServiceDisk            `json:"disk,omitempty"`
	Environment                ServiceEnvironmentType `json:"env,omitempty"`
	EnvironmentSpecificDetails any                    `json:"envSpecificDetails,omitempty"`
	NumberOfInstances          int                    `json:"numInstances,omitempty"`
	Plan                       ServicePlan            `json:"plan,omitempty"`
	PullRequestPreviewsEnabled YesNo                  `json:"pullRequestPreviewsEnabled,omitempty"`
	Region                     ServiceRegion          `json:"region,omitempty"`
}

type ServiceBackgroundWorkerDetailsPost struct {
	Disk                       ServiceDisk            `json:"disk,omitempty"`
	Environment                ServiceEnvironmentType `json:"env,omitempty"`
	EnvironmentSpecificDetails any                    `json:"envSpecificDetails,omitempty"`
	NumberOfInstances          int                    `json:"numInstances,omitempty"`
	Plan                       ServicePlan            `json:"plan,omitempty"`
	PullRequestPreviewsEnabled YesNo                  `json:"pullRequestPreviewsEnabled,omitempty"`
	Region                     ServiceRegion          `json:"region,omitempty"`
}

type ServiceCronJobDetailsPost struct {
	Environment                ServiceEnvironmentType `json:"env,omitempty"`
	EnvironmentSpecificDetails any                    `json:"envSpecificDetails,omitempty"`
	Plan                       ServicePlan            `json:"plan,omitempty"`
	Region                     ServiceRegion          `json:"region,omitempty"`
	Schedule                   string                 `json:"schedule,omitempty"`
}

type ServiceUpdateBody struct {
	AutoDeploy     YesNo          `json:"autoDeploy,omitempty"`
	Branch         string         `json:"branch,omitempty"`
	Name           string         `json:"name,omitempty"`
	ServiceDetails ServiceDetails `json:"serviceDetails,omitempty"`
}

type ParentServer struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceDetails struct {
	BuildCommand               string                              `json:"buildCommand,omitempty"`
	ParentServer               ParentServer                        `json:"parentServer,omitempty"`
	PublishPath                string                              `json:"publishPath,omitempty"`
	PullRequestPreviewsEnabled YesNo                               `json:"pullRequestPreviewsEnabled,omitempty"`
	Url                        string                              `json:"url,omitempty"`
	Env                        ServiceEnvironmentType              `json:"env,omitempty"`
	Plan                       ServicePlan                         `json:"plan,omitempty"`
	Region                     ServiceRegion                       `json:"region,omitempty"`
	Disk                       *ServiceDisk                        `json:"disk,omitempty"`
	HealthCheckPath            string                              `json:"healthCheckPath,omitempty"`
	EnvSpecificDetails         ServiceEnvironmentDockerDetailsPost `json:"envSpecificDetails,omitempty"`
}

type ServiceUpdate struct {
	Id             string         `json:"id,omitempty"`
	AutoDeploy     string         `json:"autoDeploy,omitempty"`
	Branch         string         `json:"branch,omitempty"`
	CreatedAt      time.Time      `json:"createdAt,omitempty"`
	Name           string         `json:"name,omitempty"`
	NotifyOnFail   string         `json:"notifyOnFail,omitempty"`
	OwnerId        string         `json:"ownerId,omitempty"`
	Repo           string         `json:"repo,omitempty"`
	Slug           string         `json:"slug,omitempty"`
	Suspended      string         `json:"suspended,omitempty"`
	Suspenders     []string       `json:"suspenders,omitempty"`
	Type           string         `json:"type,omitempty"`
	UpdatedAt      time.Time      `json:"updatedAt,omitempty"`
	ServiceDetails ServiceDetails `json:"serviceDetails,omitempty"`
}

type ResourceGetEnvOptions struct {
	Cursor string `url:"cursor"`
	Limit  int    `url:"limit"`
}

type EnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ServiceEnvVar struct {
	Env    EnvVar `json:"envVar"`
	Cursor string `json:"cursor,omitempty"`
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

func (s *ServicesService) CreateService(ctx context.Context, body *ServiceCreateBody) (*CreatedService, *http.Response, error) {
	url := "services"
	req, err := s.client.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var service *CreatedService
	res, err := s.client.Do(ctx, req, &service)
	if err != nil {
		return nil, nil, err
	}

	return service, res, nil
}

func (s *ServicesService) UpdateService(ctx context.Context, serviceId string, body *ServiceUpdateBody) (*ServiceUpdate, *http.Response, error) {
	url := fmt.Sprintf("services/%s", serviceId)

	req, err := s.client.NewRequest("PATCH", url, body)
	if err != nil {
		return nil, nil, err
	}

	var serviceUpdate *ServiceUpdate
	res, err := s.client.Do(ctx, req, &serviceUpdate)
	if err != nil {
		return nil, nil, err
	}

	return serviceUpdate, res, err
}

func (s *ServicesService) DeleteService(ctx context.Context, serviceId string) *http.Response {
	url := fmt.Sprintf("services/%s", serviceId)

	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil
	}

	res, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return nil
	}

	return res
}

func (s *ServicesService) ServiceGetEnvVars(ctx context.Context, serviceId string, opts *ResourceGetEnvOptions) (*[]ServiceEnvVar, *http.Response, error) {
	url := fmt.Sprintf("services/%s/env-vars", serviceId)

	// Parse options into url query parameters
	if opts != nil {
		queryString, _ := query.Values(opts)
		url += fmt.Sprintf("?%s", queryString.Encode())
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var envVars *[]ServiceEnvVar
	res, err := s.client.Do(ctx, req, &envVars)
	if err != nil {
		return nil, nil, err
	}

	return envVars, res, err
}

func (s *ServicesService) ServiceUpdateEnvVars(ctx context.Context, serviceId string, opts *[]EnvVar) (*[]ServiceEnvVar, *http.Response, error) {
	url := fmt.Sprintf("services/%s/env-vars", serviceId)

	req, err := s.client.NewRequest("PUT", url, opts)
	if err != nil {
		return nil, nil, err
	}

	var envVars *[]ServiceEnvVar
	res, err := s.client.Do(ctx, req, &envVars)
	if err != nil {
		return nil, nil, err
	}

	return envVars, res, err
}
