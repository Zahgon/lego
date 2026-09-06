package api

import (
	"context"
	"crypto"
	"net/http"
	"sync"

	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/acme/api/internal/nonces"
	"github.com/go-acme/lego/v5/acme/api/internal/secure"
	"github.com/go-acme/lego/v5/acme/api/internal/sender"
)

type service struct {
	core *Core
}

type Core struct {
	doer         *sender.Doer
	nonceManager *nonces.Manager
	directory    acme.Directory

	HTTPClient *http.Client

	privateKey crypto.Signer
	kid        string
	mu         sync.RWMutex

	common         service
	Accounts       *AccountService
	Authorizations *AuthorizationService
	Certificates   *CertificateService
	Challenges     *ChallengeService
	Orders         *OrderService
}

func New(httpClient *http.Client, userAgent, caDirURL, kid string, privateKey crypto.Signer) (*Core, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newCore(httpClient *http.Client, doer *sender.Doer, dir acme.Directory, kid string, privateKey crypto.Signer) (*Core, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Core) signer() *secure.Signer { _ = "STUB: not implemented"; return nil }

func (a *Core) setKid(kid string) { _ = "STUB: not implemented"; return }

func (a *Core) setPrivateKey(privateKey crypto.Signer) { _ = "STUB: not implemented"; return }

func (a *Core) post(ctx context.Context, uri string, reqBody, response any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Core) postAsGet(ctx context.Context, uri string, response any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Core) retrievablePost(ctx context.Context, uri string, content []byte, response any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Core) signedPost(ctx context.Context, uri string, content []byte, response any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Core) GetKeyAuthorization(token string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (a *Core) GetDirectory() acme.Directory {
	_ = "STUB: not implemented"
	return *new(acme.Directory)
}

func getDirectory(ctx context.Context, do *sender.Doer, caDirURL string) (acme.Directory, error) {
	_ = "STUB: not implemented"
	return *new(acme.Directory), nil
}
