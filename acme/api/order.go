package api

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/acme"
)

type OrderOptions struct {
	NotBefore time.Time
	NotAfter  time.Time

	Profile string

	ReplacesCertID string
}

type OrderService service

func (o *OrderService) New(ctx context.Context, domains []string, opts *OrderOptions) (acme.ExtendedOrder, error) {
	_ = "STUB: not implemented"
	return *new(acme.ExtendedOrder), nil
}

func (o *OrderService) Get(ctx context.Context, orderURL string) (acme.ExtendedOrder, error) {
	_ = "STUB: not implemented"
	return *new(acme.ExtendedOrder), nil
}

func (o *OrderService) UpdateForCSR(ctx context.Context, orderURL string, csr []byte) (acme.ExtendedOrder, error) {
	_ = "STUB: not implemented"
	return *new(acme.ExtendedOrder), nil
}
