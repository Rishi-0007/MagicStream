package utils

import (
	"context"
	"google.golang.org/api/idtoken"
)

func VerifyGoogleIDToken(ctx context.Context, token, audience string) (map[string]interface{}, error) {
	payload, err := idtoken.Validate(ctx, token, audience)
	if err != nil { return nil, err }
	return payload.Claims, nil
}
