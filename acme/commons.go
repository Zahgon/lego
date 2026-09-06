package acme

import (
	"encoding/json"
	"time"
)

const (
	StatusDeactivated = "deactivated"
	StatusExpired     = "expired"
	StatusInvalid     = "invalid"
	StatusPending     = "pending"
	StatusProcessing  = "processing"
	StatusReady       = "ready"
	StatusRevoked     = "revoked"
	StatusUnknown     = "unknown"
	StatusValid       = "valid"
)

const (
	CRLReasonUnspecified          uint = 0
	CRLReasonKeyCompromise        uint = 1
	CRLReasonCACompromise         uint = 2
	CRLReasonAffiliationChanged   uint = 3
	CRLReasonSuperseded           uint = 4
	CRLReasonCessationOfOperation uint = 5
	CRLReasonCertificateHold      uint = 6
	CRLReasonRemoveFromCRL        uint = 8
	CRLReasonPrivilegeWithdrawn   uint = 9
	CRLReasonAACompromise         uint = 10
)

type Directory struct {
	NewNonceURL   string `json:"newNonce"`
	NewAccountURL string `json:"newAccount"`
	NewOrderURL   string `json:"newOrder"`
	NewAuthzURL   string `json:"newAuthz"`
	RevokeCertURL string `json:"revokeCert"`
	KeyChangeURL  string `json:"keyChange"`
	Meta          Meta   `json:"meta"`
	RenewalInfo   string `json:"renewalInfo"`
}

type Meta struct {
	TermsOfService string `json:"termsOfService"`

	Website string `json:"website"`

	CaaIdentities []string `json:"caaIdentities"`

	ExternalAccountRequired bool `json:"externalAccountRequired"`

	Profiles map[string]string `json:"profiles"`
}

type ExtendedAccount struct {
	Account

	Location string `json:"accountURL,omitempty"`
}

type Account struct {
	Status string `json:"status,omitempty"`

	Contact []string `json:"contact,omitempty"`

	TermsOfServiceAgreed bool `json:"termsOfServiceAgreed,omitempty"`

	Orders string `json:"orders,omitempty"`

	OnlyReturnExisting bool `json:"onlyReturnExisting,omitempty"`

	ExternalAccountBinding json.RawMessage `json:"externalAccountBinding,omitempty"`
}

type ExtendedOrder struct {
	Order

	Location string `json:"-"`
}

type Order struct {
	Status string `json:"status,omitempty"`

	Expires string `json:"expires,omitempty"`

	Identifiers []Identifier `json:"identifiers"`

	Profile string `json:"profile,omitempty"`

	NotBefore string `json:"notBefore,omitempty"`

	NotAfter string `json:"notAfter,omitempty"`

	Error *ProblemDetails `json:"error,omitempty"`

	Authorizations []string `json:"authorizations,omitempty"`

	Finalize string `json:"finalize,omitempty"`

	Certificate string `json:"certificate,omitempty"`

	Replaces string `json:"replaces,omitempty"`
}

func (r *Order) Err() error { _ = "STUB: not implemented"; return nil }

type Authorization struct {
	Status string `json:"status"`

	Expires time.Time `json:"expires,omitzero"`

	Identifier Identifier `json:"identifier"`

	Challenges []Challenge `json:"challenges,omitempty"`

	Wildcard bool `json:"wildcard,omitempty"`
}

type ExtendedChallenge struct {
	Challenge

	RetryAfter time.Duration `json:"-"`

	AuthorizationURL string `json:"-"`
}

type Challenge struct {
	Type string `json:"type"`

	URL string `json:"url"`

	Status string `json:"status"`

	Validated time.Time `json:"validated,omitzero"`

	Error *ProblemDetails `json:"error,omitempty"`

	Token string `json:"token"`

	AccountURI string `json:"accounturi,omitempty"`

	IssuerDomainNames []string `json:"issuer-domain-names,omitempty"`

	KeyAuthorization string `json:"keyAuthorization"`
}

func (c *Challenge) Err() error { _ = "STUB: not implemented"; return nil }

type Identifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CSRMessage struct {
	Csr string `json:"csr"`
}

type RevokeCertMessage struct {
	Certificate string `json:"certificate"`

	Reason *uint `json:"reason,omitempty"`
}

type RawCertificate struct {
	Cert   []byte
	Issuer []byte
}

type ExtendedRenewalInfo struct {
	RenewalInfo

	RetryAfter time.Duration `json:"-"`
}

type RenewalInfo struct {
	SuggestedWindow Window `json:"suggestedWindow"`

	ExplanationURL string `json:"explanationURL"`
}

type Window struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type KeyChange struct {
	Account string          `json:"account"`
	OldKey  json.RawMessage `json:"oldKey"`
}
