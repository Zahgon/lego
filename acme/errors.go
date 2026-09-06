package acme

import (
	"time"
)

const (
	errNS = "urn:ietf:params:acme:error:"

	InvalidProfileErrorType = errNS + "invalidProfile"

	AlreadyReplacedErrorType = errNS + "alreadyReplaced"

	AccountDoesNotExistErrorType = errNS + "accountDoesNotExist"

	AlreadyRevokedErrorType = errNS + "alreadyRevoked"

	BadCSRErrorType = errNS + "badCSR"

	BadNonceErrorType = errNS + "badNonce"

	BadPublicKeyErrorType = errNS + "badPublicKey"

	BadRevocationReasonErrorType = errNS + "badRevocationReason"

	BadSignatureAlgorithmErrorType = errNS + "badSignatureAlgorithm"

	CaaErrorType = errNS + "caa"

	CompoundErrorType = errNS + "compound"

	ConnectionErrorType = errNS + "connection"

	DNSErrorType = errNS + "dns"

	ExternalAccountRequiredErrorType = errNS + "externalAccountRequired"

	IncorrectResponseErrorType = errNS + "incorrectResponse"

	InvalidContactErrorType = errNS + "invalidContact"

	MalformedErrorType = errNS + "malformed"

	OrderNotReadyErrorType = errNS + "orderNotReady"

	RateLimitedErrorType = errNS + "rateLimited"

	RejectedIdentifierErrorType = errNS + "rejectedIdentifier"

	ServerInternalErrorType = errNS + "serverInternal"

	TLSErrorType = errNS + "tls"

	UnauthorizedErrorType = errNS + "unauthorized"

	UnsupportedContactErrorType = errNS + "unsupportedContact"

	UnsupportedIdentifierErrorType = errNS + "unsupportedIdentifier"

	UserActionRequiredErrorType = errNS + "userActionRequired"
)

type ProblemDetails struct {
	Type        string       `json:"type,omitempty"`
	Detail      string       `json:"detail,omitempty"`
	HTTPStatus  int          `json:"status,omitempty"`
	Instance    string       `json:"instance,omitempty"`
	SubProblems []SubProblem `json:"subproblems,omitempty"`

	Method string `json:"method,omitempty"`
	URL    string `json:"url,omitempty"`
}

func (p *ProblemDetails) Error() string { _ = "STUB: not implemented"; return "" }

type SubProblem struct {
	Type       string     `json:"type,omitempty"`
	Detail     string     `json:"detail,omitempty"`
	Identifier Identifier `json:"identifier"`
}

type NonceError struct {
	*ProblemDetails
}

func (e *NonceError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type AlreadyReplacedError struct {
	*ProblemDetails
}

func (e *AlreadyReplacedError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type RateLimitedError struct {
	*ProblemDetails

	RetryAfter time.Duration
}

func (e *RateLimitedError) Unwrap() error { _ = "STUB: not implemented"; return nil }
