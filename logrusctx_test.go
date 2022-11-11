package logrusctx_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/0x2b3bfa0/logrusctx"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultLogger(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	logrus.SetOutput(&buf)
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})

	ctx := context.Background()

	logrusctx.Logger(ctx).Info("test")
	require.Equal(t, "level=info msg=test\n", buf.String())
}

func TestLogger(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	logger := logrus.New()
	logger.SetOutput(&buf)
	logger.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})

	ctx := logrusctx.WithLogger(context.Background(), logger)

	logrusctx.Logger(ctx).Info("test")
	require.Equal(t, "level=info msg=test\n", buf.String())
}

func TestWithFields(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	logger := logrus.New()
	logger.SetOutput(&buf)
	logger.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})

	ctx := logrusctx.WithLogger(context.Background(), logger)
	ctx = logrusctx.WithFields(ctx, logrus.Fields{"1": "one", "2": "two"})
	ctx = logrusctx.WithField(ctx, "3", "three")

	logrusctx.Info(ctx, "test")
	require.Equal(t, "level=info msg=test 1=one 2=two 3=three\n", buf.String())
}

func TestWithError(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	logger := logrus.New()
	logger.SetOutput(&buf)
	logger.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})

	ctx := logrusctx.WithLogger(context.Background(), logger)
	ctx = logrusctx.WithError(ctx, assert.AnError)

	logrusctx.Info(ctx, "test")
	require.Equal(t, "level=info msg=test error=\"assert.AnError general error for testing\"\n", buf.String())
}

func TestWithTime(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	logger := logrus.New()
	logger.SetOutput(&buf)

	ctx := logrusctx.WithLogger(context.Background(), logger)
	ctx = logrusctx.WithTime(ctx, time.Unix(1, 0))

	logrusctx.Info(ctx, "test")
	require.Equal(t, "time=\"1970-01-01T00:00:01Z\" level=info msg=test\n", buf.String())
}
