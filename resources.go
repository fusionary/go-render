package render

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ParentServer struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceDetails struct {
	BuildCommand               string       `json:"buildCommand,omitempty"`
	ParentServer               ParentServer `json:"parentServer,omitempty"`
	PublishPath                string       `json:"publishPath,omitempty"`
	PullRequestPreviewsEnabled string       `json:"pullRequestPreviewsEnabled,omitempty"`
	Url                        string       `json:"url,omitempty"`
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

func CreateResource(ctx context.Context, payload io.Reader) bool {
	url := "https://api.render.com/v1/services"

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

func UpdateService(ctx context.Context, serviceId string, payload io.Reader) (*ServiceUpdate, error) {
	url := fmt.Sprintf("https://api.render.com/v1/services/%s", serviceId)

	req, _ := http.NewRequest("PATCH", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", ctx.Value("token")))

	var serviceUpdate *ServiceUpdate
	res, err := http.DefaultClient.Do(req)
	decErr := json.NewDecoder(res.Body).Decode(&serviceUpdate)
	if decErr == io.EOF {
		decErr = nil // ignore EOF errors caused by empty response body
	}
	if decErr != nil {
		err = decErr
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		return serviceUpdate, err
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		log.Printf("%d %s", res.StatusCode, body)
		return nil, err
	}
}

func GetEnvVars(ctx context.Context, serviceId string, opts *ResourceGetEnvOptions) (*[]ServiceEnvVar, error) {
	url := fmt.Sprintf("https://api.render.com/v1/services/%s/env-vars", serviceId)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", ctx.Value("token")))

	var envVars *[]ServiceEnvVar
	res, err := http.DefaultClient.Do(req)
	decErr := json.NewDecoder(res.Body).Decode(&envVars)
	if decErr == io.EOF {
		decErr = nil // ignore EOF errors caused by empty response body
	}
	if decErr != nil {
		err = decErr
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		return envVars, err
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		log.Printf("%d %s", res.StatusCode, body)
		return nil, err
	}
}
