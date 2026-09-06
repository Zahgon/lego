package http01

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/acme/api"
	"github.com/go-acme/lego/v5/challenge"
)

const PathPrefix = "/.well-known/acme-challenge/"

type ValidateFunc func(ctx context.Context, core *api.Core, domain string, chlng acme.Challenge) error

type ChallengeOption func(*Challenge) error

func SetDelay(delay time.Duration) ChallengeOption {
	_ = "STUB: not implemented"
	return *new(ChallengeOption)
}

func ChallengePath(token string) string { _ = "STUB: not implemented"; return "" }

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

func isBase64url(s string) bool { _ = "STUB: not implemented"; return false }
