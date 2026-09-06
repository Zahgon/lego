package dnspersist01

import (
	"time"
)

const (
	policyWildcard    = "wildcard"
	paramAccountURI   = "accounturi"
	paramPolicy       = "policy"
	paramPersistUntil = "persistuntil"
)

type IssueValue struct {
	IssuerDomainName string
	AccountURI       string
	Policy           string
	PersistUntil     time.Time
}

func (v *IssueValue) match(other IssueValue) bool { _ = "STUB: not implemented"; return false }

func buildIssueValue(issuerDomainName, accountURI string, wildcard bool, persistUntil time.Time) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//nolint:gocyclo // parsing and validating tagged parameters requires branching
func parseIssueValue(value string) (*IssueValue, error) { _ = "STUB: not implemented"; return nil, nil }

func trimWSP(s string) string { _ = "STUB: not implemented"; return "" }
