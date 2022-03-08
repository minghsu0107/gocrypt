package utils

import (
	"context"
	"time"
)

// LoopTick runs a given action in a loop in periods of 't' duration. It exits
// when the context is cancelled
func LoopTick(ctx context.Context, t time.Duration, action func(errChan chan error)) error {
	ticker := time.NewTicker(t)
	defer ticker.Stop()

	errChan := make(chan error)

	for {

		// Run action
		go action(errChan)

		select {
		// Return if context is cancelled
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errChan:
			if err != nil {
				return err
			}
		// Break select every tick
		case <-ticker.C:
		}
	}
}
