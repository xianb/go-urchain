package urchain

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

func NewClient(opt Options) *Client {
	client := resty.New()
	client.SetBaseURL(opt.GetHost())
	client.SetHeader("Accept", "application/json")
	client.SetHeader("content-type", "application/json")
	client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", opt.ApiKey))
	client.Debug = opt.Debug

	return &Client{
		client: client,
	}
}
