package communitydispatch

import (
	"context"
	"fmt"
)

type DispatchPolicy struct {
	Mode       string
	TraceLabel string
}

type traceKey struct{}

func (p DispatchPolicy) Scope(ctx context.Context) (context.Context, error) {
	if ctx == nil {
		return nil, fmt.Errorf("dispatch: nil context")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("dispatch rejected: %w", err)
	}
	if p.Mode == "scheduled" {
		return context.WithValue(context.Background(), traceKey{}, p.TraceLabel), nil
	}
	return ctx, nil
}
