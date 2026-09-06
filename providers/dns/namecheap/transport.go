package namecheap

import (
	"net/http"
	"net/url"
	"sync"
)

const (
	envHTTPProxy       = "HTTP_PROXY"
	envHTTPProxyLower  = "http_proxy"
	envHTTPSProxy      = "HTTPS_PROXY"
	envHTTPSProxyLower = "https_proxy"
	envNoProxy         = "NO_PROXY"
	envNoProxyLower    = "no_proxy"
	envRequestMethod   = "REQUEST_METHOD"
)

var (
	envProxyOnce      sync.Once
	envProxyFuncValue func(*url.URL) (*url.URL, error)
)

func defaultTransport(namespace string) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

func envProxyFunc(namespace string) func(*url.URL) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil
}

func proxyFromEnvironment(namespace string) func(req *http.Request) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil
}

func getEnv(namespace, baseEnvName, baseEnvNameLower string) string {
	_ = "STUB: not implemented"
	return ""
}
