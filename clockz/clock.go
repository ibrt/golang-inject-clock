package clockz

import (
	"context"

	"github.com/benbjohnson/clock"
	"github.com/ibrt/golang-errors/errorz"
	"github.com/ibrt/golang-inject/injectz"
)

type contextKey int

const (
	clockContextKey contextKey = iota
)

var (
	_ injectz.Initializer = Initializer
)

// Clock describes a clock.
type Clock clock.Clock

// Initializer is a Clock initializer.
func Initializer(_ context.Context) (injectz.Injector, injectz.Releaser) {
	return NewSingletonInjector(clock.New()), injectz.NewNoopReleaser()
}

// NewSingletonInjector always injects the given Clock.
func NewSingletonInjector(clock Clock) injectz.Injector {
	return injectz.NewSingletonInjector(clockContextKey, clock)
}

// Get extracts the Clock from context, panics if not found.
func Get(ctx context.Context) Clock {
	clk := MaybeGet(ctx)
	errorz.Assertf(clk != nil, "clockz: not initialized", errorz.Skip())
	return clk
}

// MaybeGet is like Get but returns nil if not found.
func MaybeGet(ctx context.Context) Clock {
	if clk, ok := ctx.Value(clockContextKey).(Clock); ok {
		return clk
	}
	return nil
}
