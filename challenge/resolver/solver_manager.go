package resolver

import (
	"context"

	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/acme/api"
	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/challenge/dns01"
	"github.com/go-acme/lego/v5/challenge/dnspersist01"
	"github.com/go-acme/lego/v5/challenge/http01"
	"github.com/go-acme/lego/v5/challenge/tlsalpn01"
)

type byType []acme.Challenge

func (a byType) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a byType) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (a byType) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type SolverManager struct {
	core    *api.Core
	solvers map[challenge.Type]solver
}

func NewSolversManager(core *api.Core) *SolverManager { _ = "STUB: not implemented"; return nil }

func (c *SolverManager) SetHTTP01Provider(p challenge.Provider, opts ...http01.ChallengeOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SolverManager) SetTLSALPN01Provider(p challenge.Provider, opts ...tlsalpn01.ChallengeOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SolverManager) SetDNS01Provider(p challenge.Provider, opts ...dns01.ChallengeOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SolverManager) SetDNSPersist01(opts ...dnspersist01.ChallengeOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SolverManager) Remove(chlgType challenge.Type) { _ = "STUB: not implemented"; return }

func (c *SolverManager) ResetSolvers() { _ = "STUB: not implemented"; return }

func (c *SolverManager) chooseSolver(authz acme.Authorization) solver {
	_ = "STUB: not implemented"
	return *new(solver)
}

func validate(ctx context.Context, core *api.Core, domain string, chlg acme.Challenge) error {
	_ = "STUB: not implemented"
	return nil
}

func checkChallengeStatus(chlng acme.ExtendedChallenge) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func checkAuthorizationStatus(authz acme.Authorization) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func solversToString(s map[challenge.Type]solver) string { _ = "STUB: not implemented"; return "" }
