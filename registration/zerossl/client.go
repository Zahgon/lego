package zerossl

import (
	"context"
	"net/http"
	"net/url"
)

const EnvZeroSSLAccessKey = "ZERO_SSL_ACCESS_KEY"

const defaultBaseURL = "https://api.zerossl.com"

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

func NewClient() *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) GenerateEAB(ctx context.Context, accessKey string) (*APIResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) GenerateEABFromEmail(ctx context.Context, email string) (*APIResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) do(req *http.Request) (*APIResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseError(req *http.Request, resp *http.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func IsZeroSSL(server string) bool { _ = "STUB: not implemented"; return false }
