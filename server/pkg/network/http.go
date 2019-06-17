package network

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	client http.Client
}

func NewClient() HttpClient {
	return HttpClient{
		http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *HttpClient) Get(ctx context.Context, url string, target interface{}) (int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	rsp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}

	if target == nil {
		return rsp.StatusCode, nil
	}

	bodyBytes, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return 0, err
	}

	return rsp.StatusCode, json.Unmarshal(bodyBytes, target)
}
