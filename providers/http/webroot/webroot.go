package webroot

import (
	"context"
)

type HTTPProvider struct {
	path string
}

func NewHTTPProvider(path string) (*HTTPProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *HTTPProvider) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *HTTPProvider) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}
