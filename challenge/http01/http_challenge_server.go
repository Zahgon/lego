package http01

import (
	"context"
	"io/fs"
	"net"

	"github.com/go-acme/lego/v5/challenge"
)

var _ challenge.Provider = (*ProviderServer)(nil)

type Options struct {
	Network         string
	NetworkStack    challenge.NetworkStack
	Address         string
	SocketMode      fs.FileMode
	ProxyHeaderName string
}

type ProviderServer struct {
	network string
	address string

	socketMode fs.FileMode

	matcher  domainMatcher
	done     chan bool
	listener net.Listener
}

func NewProviderServerWithOptions(opts Options) *ProviderServer {
	_ = "STUB: not implemented"
	return nil
}

func NewProviderServer(host, port string) *ProviderServer { _ = "STUB: not implemented"; return nil }

func NewUnixProviderServer(socketPath string, socketMode fs.FileMode) *ProviderServer {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProviderServer) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProviderServer) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProviderServer) GetAddress() string { _ = "STUB: not implemented"; return "" }

func getMatcher(proxyHeaderName string) domainMatcher {
	_ = "STUB: not implemented"
	return *new(domainMatcher)
}

func (s *ProviderServer) serve(domain, token, keyAuth string) { _ = "STUB: not implemented"; return }
