# golang-inject-clock
[![Go Reference](https://pkg.go.dev/badge/github.com/ibrt/golang-inject-clock.svg)](https://pkg.go.dev/github.com/ibrt/golang-inject-clock)
![CI](https://github.com/ibrt/golang-inject-clock/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/gh/ibrt/golang-inject-clock/branch/main/graph/badge.svg?token=BQVP881F9Z)](https://codecov.io/gh/ibrt/golang-inject-clock)

Clock module for the [golang-inject](https://github.com/ibrt/golang-inject) framework.

### Basic Usage

This module injects a "clock" into Go context. It provides both a real and a mock implementation for use in tests. 
Beyond being useful by itself, it is also a minimal example of how to tie together modules using the
[golang-inject](https://github.com/ibrt/golang-inject) framework, and how to easily test implementations using the
[golang-fixtures](https://github.com/ibrt/golang-fixtures) test suites.

```go
// main.go

package main

import (
    "net/http"
    "time"

    "github.com/ibrt/golang-inject/injectz"
    "github.com/ibrt/golang-inject-clock/clockz"
)

func main() {
    injector, releaser := injectz.Initialize(clockz.Initializer)
    defer releaser()

    middleware := injectz.NewMiddleware(injector)
    mux := http.NewServeMux()
    mux.Handle("/", middleware(http.HandlerFunc(handler)))
    _ = http.ListenAndServe(":3000", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
    _, _ = w.Write([]byte(clockz.Get(r.Context()).Now().Format(time.RFC3339Nano)))
}
```

```go
// main_test.go

package main

import (
    "context"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/ibrt/golang-fixtures/fixturez"
    "github.com/ibrt/golang-inject-clock/clockz/testclockz"
    "github.com/stretchr/testify/require"
)

var (
    _ fixturez.Suite = &Suite{}
)

type Suite struct {
    *fixturez.DefaultConfigMixin
    Clockz *testclockz.MockHelper
}

func TestSuite(t *testing.T) {
    fixturez.RunSuite(t, &Suite{})
}

func (s *Suite) TestHandler(ctx context.Context, t *testing.T) {
    const nowStr = "2009-11-18T08:04:34.829482969Z"
    now, err := time.Parse(time.RFC3339Nano, nowStr)
    require.NoError(t, err)
    s.Clockz.Clock.Set(now)

    w := httptest.NewRecorder()
    r := httptest.NewRequest("GET", "/", nil).WithContext(ctx)

    handler(w, r)

    require.Equal(t, http.StatusOK, w.Code)
    require.Equal(t, nowStr, w.Body.String())
}
```

### Developers

Contributions are welcome, please check in on proposed implementation before sending a PR. You can validate your changes
using the `./test.sh` script.