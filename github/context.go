package github

import (
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/net/context"
)

// unexported key type prevents collisions
type key int

const (
	userKey key = iota
)

// WithUser returns a copy of ctx that stores the Github User.
func WithUser(ctx context.Context, user *github.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// UserFromContext returns the Github User from the ctx.
func UserFromContext(ctx context.Context) (*github.User, error) {
	user, ok := ctx.Value(userKey).(*github.User)
	if !ok {
		return nil, fmt.Errorf("Context missing Github User")
	}
	return user, nil
}
