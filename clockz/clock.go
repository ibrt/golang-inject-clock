package clockz

import (
	"context"

	clocklib "github.com/benbjohnson/clock"
	"github.com/ibrt/golang-inject/injectz"
)

type contextKey int

const (
	clockContextKey contextKey = iota
)

var (
	_ injectz.Initializer = Initializer

	defaultClock = clocklib.New()
)

// Clock describes a clock.
type Clock clocklib.Clock

// Initializer is a Clock initializer.
func Initializer(_ context.Context) (injectz.Injector, injectz.Releaser) {
	return injectz.NewNoopInjector(), injectz.NewNoopReleaser()
}

// NewSingletonInjector always injects the given Clock.
func NewSingletonInjector(clock Clock) injectz.Injector {
	return injectz.NewSingletonInjector(clockContextKey, clock)
}

// Get extracts the Clock from context, returns the default clock if not found.
func Get(ctx context.Context) Clock {
	if clock, ok := ctx.Value(clockContextKey).(Clock); ok {
		return clock
	}
	return defaultClock
}
