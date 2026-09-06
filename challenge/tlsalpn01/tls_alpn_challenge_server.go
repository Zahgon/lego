package tlsalpn01

import (
	"context"
	"net"

	"github.com/go-acme/lego/v5/challenge"
)

const (
	ACMETLS1Protocol = "acme-tls/1"

	defaultTLSPort = "443"
)

var _ challenge.Provider = (*ProviderServer)(nil)

type Options struct {
	Network      string
	NetworkStack challenge.NetworkStack
	Host         string
	Port         string
}

type ProviderServer struct {
	network string
	address string

	listener net.Listener
}

func NewProviderServerWithOptions(opts Options) *ProviderServer {
	_ = "STUB: not implemented"
	return nil
}

func NewProviderServer(host, port string) *ProviderServer { _ = "STUB: not implemented"; return nil }

func (s *ProviderServer) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProviderServer) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProviderServer) GetAddress() string { _ = "STUB: not implemented"; return "" }
