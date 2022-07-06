package render

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

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
