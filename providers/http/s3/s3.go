package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type HTTPProvider struct {
	bucket string
	client *s3.Client
}

func NewHTTPProvider(bucket string) (*HTTPProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *HTTPProvider) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *HTTPProvider) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}
