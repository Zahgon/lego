package dns01

import (
	"context"
)

type PreCheckFunc func(ctx context.Context, fqdn, value string) (bool, error)

type WrapPreCheckFunc func(ctx context.Context, domain, fqdn, value string, check PreCheckFunc) (bool, error)

func WrapPreCheck(wrap WrapPreCheckFunc) ChallengeOption {
	_ = "STUB: not implemented"
	return *new(ChallengeOption)
}

type preCheck struct {
	checkFunc WrapPreCheckFunc

	requireAuthoritativeNssPropagation bool

	requireRecursiveNssPropagation bool
}

func newPreCheck() preCheck { _ = "STUB: not implemented"; return *new(preCheck) }

func (p preCheck) call(ctx context.Context, domain, fqdn, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p preCheck) checkDNSPropagation(ctx context.Context, fqdn, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
