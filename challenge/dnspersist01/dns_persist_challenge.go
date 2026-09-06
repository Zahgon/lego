package dnspersist01

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/acme/api"
	"github.com/go-acme/lego/v5/challenge"
)

const validationLabel = "_validation-persist"

type ValidateFunc func(ctx context.Context, core *api.Core, domain string, chlng acme.Challenge) error

type Challenge struct {
	core     *api.Core
	validate ValidateFunc
	provider challenge.PersistentProvider
	preCheck preCheck

	userSuppliedIssuerDomainName string
	persistUntil                 time.Time
}

func NewChallenge(core *api.Core, validate ValidateFunc, provider challenge.PersistentProvider, opts ...ChallengeOption) (*Challenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Challenge) Solve(ctx context.Context, authz acme.Authorization) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Challenge) waitForPropagation(ctx context.Context, domain, fqdn string, matcher RecordMatcher) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Challenge) selectIssuerDomainName(challIssuers []string, records []TXTRecord, accountURI string, wildcard bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Challenge) hasMatchingRecord(records []TXTRecord, issuerDomainName, accountURI string, wildcard bool) bool {
	_ = "STUB: not implemented"
	return false
}

type ChallengeInfo struct {
	FQDN string

	Value string

	IssuerDomainName string
}

func GetChallengeInfo(authz acme.Authorization, issuerDomainName, accountURI string, persistUntil time.Time) (ChallengeInfo, error) {
	_ = "STUB: not implemented"
	return *new(ChallengeInfo), nil
}

func getValidationDomainName(domain string) string { _ = "STUB: not implemented"; return "" }
