package testclockz

import (
	"context"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/ibrt/golang-fixtures/fixturez"
	"github.com/ibrt/golang-inject/injectz"

	"github.com/ibrt/golang-inject-clock/clockz"
)

var (
	_ fixturez.BeforeSuite = &Helper{}
	_ fixturez.AfterSuite  = &Helper{}
	_ fixturez.BeforeTest  = &MockHelper{}
	_ fixturez.AfterTest   = &MockHelper{}
)

// Helper is a test helper for Clock.
type Helper struct {
	releaser injectz.Releaser
}

// BeforeSuite implements fixturez.BeforeSuite.
func (f *Helper) BeforeSuite(ctx context.Context, _ *testing.T) context.Context {
	injector, releaser := clockz.Initializer(ctx)
	f.releaser = releaser
	return injector(ctx)
}

// AfterSuite implements fixturez.AfterSuite.
func (f *Helper) AfterSuite(_ context.Context, _ *testing.T) {
	f.releaser()
	f.releaser = nil
}

// MockHelper is a test helper for Clock.
type MockHelper struct {
	Clock *clock.Mock
}

// BeforeTest implements fixturez.BeforeTest.
func (f *MockHelper) BeforeTest(ctx context.Context, _ *testing.T) context.Context {
	f.Clock = clock.NewMock()
	f.Clock = clock.NewMock()
	f.Clock.Set(time.Now().UTC())
	return clockz.NewSingletonInjector(f.Clock)(ctx)
}

// AfterTest implements fixturez.AfterTest.
func (f *MockHelper) AfterTest(_ context.Context, _ *testing.T) {
	f.Clock = nil
}
