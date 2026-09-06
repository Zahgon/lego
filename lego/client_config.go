package lego

import (
	"crypto/x509"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/registration"
)

const (
	caCertificatesEnvVar = "LEGO_CA_CERTIFICATES"

	caSystemCertPool = "LEGO_CA_SYSTEM_CERT_POOL"

	caServerNameEnvVar = "LEGO_CA_SERVER_NAME"
)

type Config struct {
	CADirURL    string
	User        registration.User
	UserAgent   string
	HTTPClient  *http.Client
	Certificate CertificateConfig
}

func NewConfig(user registration.User) *Config { _ = "STUB: not implemented"; return nil }

type CertificateConfig struct {
	Timeout             time.Duration
	OverallRequestLimit int
}

func createDefaultHTTPClient() *http.Client { _ = "STUB: not implemented"; return nil }

func initCertPool() *x509.CertPool { _ = "STUB: not implemented"; return nil }

func CreateCertPool(caCerts []string, useSystemCertPool bool) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newCertPool(useSystemCertPool bool) *x509.CertPool { _ = "STUB: not implemented"; return nil }
