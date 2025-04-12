package util

import (
	"resty.dev/v3"
)

// RestClient is a wrapper around the resty client

type RestClient struct {
	*resty.Client
}

// NewRestClient creates a new rest client
func NewRestClient() *RestClient {
	r := resty.New()
	r.SetRetryCount(3)
	return &RestClient{resty.New()}
}
