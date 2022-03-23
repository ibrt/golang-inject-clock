package clockz_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ibrt/golang-inject-clock/clockz"
)

func TestModule(t *testing.T) {
	injector, releaser := clockz.Initializer(context.Background())
	defer releaser()

	clock := clockz.Get(context.Background())
	require.NotNil(t, clock)

	clock = clockz.Get(injector(context.Background()))
	require.NotNil(t, clock)

	clock = clockz.Get(clockz.NewSingletonInjector(clock)(context.Background()))
	require.NotNil(t, clock)
}
