package lego

import (
	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/acme/api"
	"github.com/go-acme/lego/v5/certificate"
	"github.com/go-acme/lego/v5/challenge/resolver"
	"github.com/go-acme/lego/v5/registration"
)

type Client struct {
	Certificate  *certificate.Certifier
	Challenge    *resolver.SolverManager
	Registration *registration.Registrar
	core         *api.Core
}

func NewClient(config *Config) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Client) GetServerMetadata() acme.Meta { _ = "STUB: not implemented"; return *new(acme.Meta) }
