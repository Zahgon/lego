package resolver

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/internal/errutils"
)

type solver interface {
	Solve(ctx context.Context, authorization acme.Authorization) error
}

type preSolver interface {
	PreSolve(ctx context.Context, authorization acme.Authorization) error
}

type cleanup interface {
	CleanUp(ctx context.Context, authorization acme.Authorization) error
}

type sequential interface {
	Sequential() (bool, time.Duration)
}

type selectedAuthSolver struct {
	authz  acme.Authorization
	solver solver
}

type Prober struct {
	solverManager *SolverManager
}

func NewProber(solverManager *SolverManager) *Prober { _ = "STUB: not implemented"; return nil }

func (p *Prober) Solve(ctx context.Context, authorizations []acme.Authorization) error {
	_ = "STUB: not implemented"
	return nil
}

func sequentialSolve(ctx context.Context, authSolvers []*selectedAuthSolver, failures *errutils.DomainsError) {
	_ = "STUB: not implemented"
	return
}

func parallelSolve(ctx context.Context, authSolvers []*selectedAuthSolver, failures *errutils.DomainsError) {
	_ = "STUB: not implemented"
	return
}

func cleanUp(ctx context.Context, solvr solver, authz acme.Authorization) {
	_ = "STUB: not implemented"
	return
}
