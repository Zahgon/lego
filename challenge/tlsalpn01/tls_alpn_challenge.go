package tlsalpn01

import (
	"context"
	"crypto/tls"
	"encoding/asn1"
	"time"

	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/acme/api"
	"github.com/go-acme/lego/v5/challenge"
)

var idPeAcmeIdentifierV1 = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 1, 31}

type ValidateFunc func(ctx context.Context, core *api.Core, domain string, chlng acme.Challenge) error

type ChallengeOption func(*Challenge) error

func SetDelay(delay time.Duration) ChallengeOption {
	_ = "STUB: not implemented"
	return *new(ChallengeOption)
}

type Challenge struct {
	core     *api.Core
	validate ValidateFunc
	provider challenge.Provider
	delay    time.Duration
}

func NewChallenge(core *api.Core, validate ValidateFunc, provider challenge.Provider, opts ...ChallengeOption) *Challenge {
	_ = "STUB: not implemented"
	return nil
}

func (c *Challenge) Solve(ctx context.Context, authz acme.Authorization) error {
	_ = "STUB: not implemented"
	return nil
}

func ChallengeBlocks(domain, keyAuth string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ChallengeCert(domain, keyAuth string) (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
