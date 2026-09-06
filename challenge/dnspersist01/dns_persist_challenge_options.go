package dnspersist01

import (
	"time"

	"github.com/go-acme/lego/v5/challenge/internal"
)

type ChallengeOption = internal.ChallengeOption[*Challenge]

func CondOptions(condition bool, opt ...ChallengeOption) ChallengeOption {
	_ = "STUB: not implemented"
	return *new(ChallengeOption)
}

func LazyCondOption(condition bool, fn func() ChallengeOption) ChallengeOption {
	_ = "STUB: not implemented"
	return *new(ChallengeOption)
}

func CombineOptions(opts ...ChallengeOption) ChallengeOption {
	_ = "STUB: not implemented"
	return *new(ChallengeOption)
}

func WithIssuerDomainName(issuerDomainName string) ChallengeOption {
	_ = "STUB: not implemented"
	return *new(ChallengeOption)
}

func WithPersistUntil(persistUntil time.Time) ChallengeOption {
	_ = "STUB: not implemented"
	return *new(ChallengeOption)
}

func DisableAuthoritativeNssPropagationRequirement() ChallengeOption {
	_ = "STUB: not implemented"
	return *new(ChallengeOption)
}

func DisableRecursiveNSsPropagationRequirement() ChallengeOption {
	_ = "STUB: not implemented"
	return *new(ChallengeOption)
}

func PropagationWait(wait time.Duration, skipCheck bool) ChallengeOption {
	_ = "STUB: not implemented"
	return *new(ChallengeOption)
}
