# `logrusctx` [![reference][reference-badge]][reference] [![tests][tests-badge]][tests]

Port of [github.com/juju/zaputil/zapctx](https://pkg.go.dev/github.com/juju/zaputil/zapctx) for [github.com/sirupsen/logrus](https://pkg.go.dev/github.com/sirupsen/logrus)


## Example

```go
package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/0x2b3bfa0/logrusctx"
)

func main() {
	ctx := context.Background()
	logger := logrus.WithField("key", "value")

	ctx = logrusctx.WithLogger(ctx, logger)
	worker(ctx)
}

func worker(ctx context.Context) {
	logrusctx.Info(ctx, "test")
}
```

[reference-badge]: https://pkg.go.dev/badge/github.com/0x2b3bfa0/logrusctx.svg
[reference]: https://pkg.go.dev/github.com/0x2b3bfa0/logrusctx
[tests-badge]: https://github.com/0x2b3bfa0/logrusctx/actions/workflows/test.yml/badge.svg?branch=main
[tests]: https://github.com/0x2b3bfa0/logrusctx/actions/workflows/test.yml
