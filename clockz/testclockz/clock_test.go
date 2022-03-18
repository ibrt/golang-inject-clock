package testclockz_test

import (
	"context"
	"testing"
	"time"

	"github.com/ibrt/golang-fixtures/fixturez"
	"github.com/stretchr/testify/require"

	"github.com/ibrt/golang-inject-clock/clockz"
	"github.com/ibrt/golang-inject-clock/clockz/testclockz"
)

func TestHelpers(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
	fixturez.RunSuite(t, &MockSuite{})
}

type Suite struct {
	*fixturez.DefaultConfigMixin
	Clock *testclockz.Helper
}

func (s *Suite) TestHelper(ctx context.Context, t *testing.T) {
	clk := clockz.Get(ctx)
	require.NotNil(t, clk)
	require.NotZero(t, clk.Now())
}

type MockSuite struct {
	*fixturez.DefaultConfigMixin
	Clock *testclockz.MockHelper
}

func (s *MockSuite) TestMockHelper(ctx context.Context, t *testing.T) {
	now := time.Now().Add(-time.Minute)
	s.Clock.Mock.Set(now)
	require.Equal(t, now, clockz.Get(ctx).Now())
}
