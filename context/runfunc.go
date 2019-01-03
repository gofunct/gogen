package context

import (
	"context"
)

type RunFunc func(ctx context.Context, program string, args ...string) error
