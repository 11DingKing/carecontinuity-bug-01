package communitydispatch

import (
	"context"
	"fmt"
)

type Coordinator struct{ policy DispatchPolicy }

func NewCoordinator() *Coordinator {
	return &Coordinator{policy: DispatchPolicy{Mode: "scheduled", TraceLabel: "field-team"}}
}

func (c *Coordinator) Run(ctx context.Context, operation func(context.Context) error) error {
	callCtx, err := c.policy.Scope(ctx)
	if err != nil {
		return err
	}
	if err := operation(callCtx); err != nil {
		return fmt.Errorf("community testing cancellation operation: %w", err)
	}
	return nil
}
