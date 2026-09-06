package challenge

import (
	"context"
	"time"
)

type Provider interface {
	Present(ctx context.Context, domain, token, keyAuth string) error
	CleanUp(ctx context.Context, domain, token, keyAuth string) error
}

type ProviderTimeout interface {
	Provider
	Timeout() (timeout, interval time.Duration)
}

type PersistentProvider interface {
	Persist(ctx context.Context, fqdn, value string) error
	Timeout() (timeout, interval time.Duration)
}
