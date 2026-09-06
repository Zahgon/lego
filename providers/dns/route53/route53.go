package route53

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	awstypes "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/go-acme/lego/v5/challenge"
)

const (
	envNamespace = "AWS_"

	EnvAccessKeyID     = envNamespace + "ACCESS_KEY_ID"
	EnvSecretAccessKey = envNamespace + "SECRET_ACCESS_KEY"
	EnvRegion          = envNamespace + "REGION"
	EnvHostedZoneID    = envNamespace + "HOSTED_ZONE_ID"
	EnvMaxRetries      = envNamespace + "MAX_RETRIES"
	EnvAssumeRoleArn   = envNamespace + "ASSUME_ROLE_ARN"
	EnvExternalID      = envNamespace + "EXTERNAL_ID"
	EnvPrivateZone     = envNamespace + "PRIVATE_ZONE"

	EnvWaitForRecordSetsChanged = envNamespace + "WAIT_FOR_RECORD_SETS_CHANGED"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Region          string

	HostedZoneID  string
	MaxRetries    int
	AssumeRoleArn string
	ExternalID    string
	PrivateZone   bool

	WaitForRecordSetsChanged bool

	TTL                int
	PropagationTimeout time.Duration
	PollingInterval    time.Duration

	Client *route53.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	client *route53.Client
	config *Config
}

func NewDNSProvider() (*DNSProvider, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) Timeout() (timeout, interval time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}

func (d *DNSProvider) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) changeRecord(ctx context.Context, action awstypes.ChangeAction, hostedZoneID string, recordSet *awstypes.ResourceRecordSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) getExistingRecordSets(ctx context.Context, hostedZoneID, fqdn string) ([]awstypes.ResourceRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) getHostedZoneID(ctx context.Context, fqdn string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func createAWSConfig(ctx context.Context, config *Config) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}

func createAWSConfigCheckParams(config *Config) error { _ = "STUB: not implemented"; return nil }
