package clockz_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ibrt/golang-inject-clock/clockz"
)

func TestClock(t *testing.T) {
	injector, releaser := clockz.Initializer(context.Background())
	defer releaser()
	ctx := injector(context.Background())
	now := clockz.Get(ctx)
	require.NotNil(t, now)
	require.NotZero(t, now)
}
