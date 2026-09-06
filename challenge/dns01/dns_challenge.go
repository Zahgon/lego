package dns01

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/acme/api"
	"github.com/go-acme/lego/v5/challenge"
)

const (
	DefaultPropagationTimeout = 60 * time.Second

	DefaultPollingInterval = 2 * time.Second

	DefaultTTL = 120
)

const challengeLabel = "_acme-challenge"

type ValidateFunc func(ctx context.Context, core *api.Core, domain string, chlng acme.Challenge) error

type Challenge struct {
	core     *api.Core
	validate ValidateFunc
	provider challenge.Provider
	preCheck preCheck
}

func NewChallenge(core *api.Core, validate ValidateFunc, provider challenge.Provider, opts ...ChallengeOption) *Challenge {
	_ = "STUB: not implemented"
	return nil
}

func (c *Challenge) PreSolve(ctx context.Context, authz acme.Authorization) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Challenge) Solve(ctx context.Context, authz acme.Authorization) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Challenge) CleanUp(ctx context.Context, authz acme.Authorization) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Challenge) Sequential() (bool, time.Duration) {
	_ = "STUB: not implemented"
	return false, *new(time.Duration)
}

type sequential interface {
	Sequential() time.Duration
}

type ChallengeInfo struct {
	FQDN string

	EffectiveFQDN string

	Value string

	Prefix string
}

func (c ChallengeInfo) Domain() string { _ = "STUB: not implemented"; return "" }

func (c ChallengeInfo) EffectiveDomain() string { _ = "STUB: not implemented"; return "" }

func GetChallengeInfo(ctx context.Context, domain, keyAuth string) ChallengeInfo {
	_ = "STUB: not implemented"
	return *new(ChallengeInfo)
}

func getChallengeFQDN(ctx context.Context, fqdn string, followCNAME bool) string {
	_ = "STUB: not implemented"
	return ""
}

func getAuthorizationDomainName(domain string) string { _ = "STUB: not implemented"; return "" }
