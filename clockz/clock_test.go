package clockz_test

import (
	"context"
	"testing"

	"github.com/ibrt/golang-fixtures/fixturez"
	"github.com/stretchr/testify/require"

	"github.com/ibrt/golang-inject-clock/clockz"
)

func TestModule(t *testing.T) {
	injector, releaser := clockz.Initializer(context.Background())
	defer releaser()
	ctx := injector(context.Background())
	now := clockz.Get(ctx)
	require.NotNil(t, now)
	require.NotZero(t, now)
	require.Nil(t, clockz.MaybeGet(context.Background()))
	fixturez.RequirePanicsWith(t, "clock: not initialized", func() {
		clockz.Get(context.Background())
	})
}
