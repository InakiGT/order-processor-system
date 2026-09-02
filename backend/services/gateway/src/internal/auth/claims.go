package auth

import (
	"context"
	"errors"
)

type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c *CustomClaims) Validate(ctx context.Context) error {
	if c.Scope == "" {
		return errors.New("scope is required")
	}

	return nil
}
