package render

import (
	"context"
)

func StoreAuthentication(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, "token", token)
}
