package auth

import (
	"fmt"
	"net/url"

	"github.com/auth0/go-jwt-middleware/v3/jwks"
	"github.com/auth0/go-jwt-middleware/v3/validator"
)

func NewValidator(domain, audience string) (*validator.Validator, error) {
	issuerURL, err := url.Parse("https://" + domain + "/")
	if err != nil {
		return nil, fmt.Errorf("failed to parse issuer URL: %w", err)
	}

	provider, err := jwks.NewCachingProvider(
		jwks.WithIssuerURL(issuerURL),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create JWKS provider: %w", err)
	}

	jwtValidator, err := validator.New(
		validator.WithKeyFunc(provider.KeyFunc),
		validator.WithAlgorithm(validator.RS256),
		validator.WithIssuer(issuerURL.String()),
		validator.WithAudience(audience),
		validator.WithCustomClaims(func() *CustomClaims {
			return &CustomClaims{}
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create JWT validator: %w", err)
	}

	return jwtValidator, nil
}
