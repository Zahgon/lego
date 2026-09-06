package memcached

import (
	"context"
)

type HTTPProvider struct {
	hosts []string
}

func NewMemcachedProvider(hosts []string) (*HTTPProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *HTTPProvider) Present(_ context.Context, _, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *HTTPProvider) CleanUp(_ context.Context, _, _, _ string) error {
	_ = "STUB: not implemented"
	return nil
}
