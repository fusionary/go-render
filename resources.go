package render

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type ResourceGetEnvOptions struct {
	Cursor string `url:"cursor"`
	Limit  string `url:"limit"`
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
