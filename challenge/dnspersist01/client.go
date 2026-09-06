package dnspersist01

import (
	"sync/atomic"

	"github.com/go-acme/lego/v5/challenge/internal"
)

var defaultClient atomic.Pointer[Client]

func init() {
	defaultClient.Store(NewClient(nil))
}

func DefaultClient() *Client { _ = "STUB: not implemented"; return nil }

func SetDefaultClient(c *Client) { _ = "STUB: not implemented"; return }

type Options = internal.Options

func NewOptions() *Options { _ = "STUB: not implemented"; return nil }

type Client struct {
	core *internal.Client

	authoritativeNSPort string
}

func NewClient(opts *Options) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) ClearFqdnCache() { _ = "STUB: not implemented"; return }
