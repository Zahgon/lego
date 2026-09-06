package cmd

import (
	"context"

	"github.com/urfave/cli/v3"
)

type ListCertificate struct {
	Name           string   `json:"name,omitempty"`
	Domains        []string `json:"domains,omitempty"`
	IPs            []string `json:"ips,omitempty"`
	ExpirationDate string   `json:"expirationDate,omitempty"`
	Expired        bool     `json:"expired"`
	Issuer         string   `json:"issuer,omitempty"`
	Path           string   `json:"path,omitempty"`
}

func createListCertificates() *cli.Command { _ = "STUB: not implemented"; return nil }

func listCertificates(_ context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func listCertificatesText(basePath string) error { _ = "STUB: not implemented"; return nil }

func listCertificatesJSON(basePath string) error { _ = "STUB: not implemented"; return nil }

func readCertificates(basePath string) ([]ListCertificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
