package joker

import (
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
)

const (
	envNamespace = "JOKER_"

	EnvAPIKey   = envNamespace + "API_KEY"
	EnvUsername = envNamespace + "USERNAME"
	EnvPassword = envNamespace + "PASSWORD"
	EnvMode     = envNamespace + "API_MODE"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvSequenceInterval   = envNamespace + "SEQUENCE_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const (
	modeDMAPI = "DMAPI"
	modeSVC   = "SVC"
)

type Config struct {
	APIKey             string
	Username           string
	Password           string
	APIMode            string
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	SequenceInterval   time.Duration
	TTL                int
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

func NewDNSProvider() (challenge.ProviderTimeout, error) {
	_ = "STUB: not implemented"
	return *new(challenge.ProviderTimeout), nil
}

func NewDNSProviderConfig(config *Config) (challenge.ProviderTimeout, error) {
	_ = "STUB: not implemented"
	return *new(challenge.ProviderTimeout), nil
}
