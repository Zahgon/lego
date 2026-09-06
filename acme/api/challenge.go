package api

import (
	"context"

	"github.com/go-acme/lego/v5/acme"
)

type ChallengeService service

func (c *ChallengeService) New(ctx context.Context, chlgURL string) (acme.ExtendedChallenge, error) {
	_ = "STUB: not implemented"
	return *new(acme.ExtendedChallenge), nil
}

func (c *ChallengeService) Get(ctx context.Context, chlgURL string) (acme.ExtendedChallenge, error) {
	_ = "STUB: not implemented"
	return *new(acme.ExtendedChallenge), nil
}
