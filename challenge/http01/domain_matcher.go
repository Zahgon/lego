package http01

import (
	"net/http"
)

type domainMatcher interface {
	matches(request *http.Request, domain string) bool

	name() string
}

type hostMatcher struct{}

func (m *hostMatcher) name() string { _ = "STUB: not implemented"; return "" }

func (m *hostMatcher) matches(r *http.Request, domain string) bool {
	_ = "STUB: not implemented"
	return false
}

type arbitraryMatcher string

func (m arbitraryMatcher) name() string { _ = "STUB: not implemented"; return "" }

func (m arbitraryMatcher) matches(r *http.Request, domain string) bool {
	_ = "STUB: not implemented"
	return false
}

type forwardedMatcher struct{}

func (m *forwardedMatcher) name() string { _ = "STUB: not implemented"; return "" }

func (m *forwardedMatcher) matches(r *http.Request, domain string) bool {
	_ = "STUB: not implemented"
	return false
}

func parseForwardedHeader(s string) (elements []map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tchar(r rune) bool { _ = "STUB: not implemented"; return false }

func skipWS(s string, i int) int { _ = "STUB: not implemented"; return 0 }

func isWS(r rune) bool { _ = "STUB: not implemented"; return false }

func matchDomain(src, domain string) bool { _ = "STUB: not implemented"; return false }
