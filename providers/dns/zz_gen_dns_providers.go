package dns

import (
	"github.com/go-acme/lego/v5/challenge"
)

func NewDNSChallengeProviderByName(name string) (challenge.Provider, error) {
	_ = "STUB: not implemented"
	return *new(challenge.Provider), nil
}
