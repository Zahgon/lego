package cloudflare

import (
	"context"
	"sync"

	"github.com/go-acme/lego/v5/providers/dns/cloudflare/internal"
)

type metaClient struct {
	clientEdit *internal.Client
	clientRead *internal.Client

	zones   map[string]string
	zonesMu *sync.RWMutex
}

func newClient(config *Config) (*metaClient, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *metaClient) CreateDNSRecord(ctx context.Context, zoneID string, rr internal.Record) (*internal.Record, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *metaClient) DeleteDNSRecord(ctx context.Context, zoneID, recordID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *metaClient) ZoneIDByName(ctx context.Context, fdqn string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func extractZoneID(res []internal.Zone) (string, error) { _ = "STUB: not implemented"; return "", nil }
