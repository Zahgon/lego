package oraclecloud

import (
	"crypto/rsa"

	"github.com/nrdcg/oci-go-sdk/common/v1065"
)

type environmentConfigurationProvider struct {
	values map[string]string
}

func newEnvironmentConfigurationProvider() (*environmentConfigurationProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *environmentConfigurationProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *environmentConfigurationProvider) KeyID() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *environmentConfigurationProvider) TenancyOCID() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *environmentConfigurationProvider) UserOCID() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *environmentConfigurationProvider) KeyFingerprint() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *environmentConfigurationProvider) Region() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *environmentConfigurationProvider) AuthType() (common.AuthConfig, error) {
	_ = "STUB: not implemented"
	return *new(common.AuthConfig), nil
}

func (p *environmentConfigurationProvider) privateKeyPassword() string {
	_ = "STUB: not implemented"
	return ""
}

func getPrivateKey() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func getEnvWithStrictFallback(keys ...string) string { _ = "STUB: not implemented"; return "" }

func getEnvFileWithStrictFallback(keys ...string) []byte { _ = "STUB: not implemented"; return nil }
