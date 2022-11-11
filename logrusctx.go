package logrusctx

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

// Default holds the logger entry returned by Entry when there is no logger in
// the context.
var Default = logrus.StandardLogger()

// loggerKey holds the context key used for loggers.
type loggerKey struct{}

// WithLogger returns a new context derived from ctx that is associated with
// the given logger.
func WithLogger(ctx context.Context, logger ExtendedFieldLogger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// Logger returns the logger associated with the given context. If there is no
// logger, it will return Default.
func Logger(ctx context.Context) ExtendedFieldLogger {
	if ctx == nil {
		panic("nil context")
	}

	if logger, _ := ctx.Value(loggerKey{}).(ExtendedFieldLogger); logger != nil {
		return logger
	}

	return Default
}

// ExtendedFieldLogger generalizes the Entry and Logger types.
type ExtendedFieldLogger interface {
	WithContext(ctx context.Context) *logrus.Entry
	WithTime(t time.Time) *logrus.Entry
	logrus.FieldLogger
}

// WithField returns a new context derived from ctx that has a logger that
// always logs the given field.
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	return WithLogger(ctx, Logger(ctx).WithField(key, value))
}

// WithFields returns a new context derived from ctx that has a logger that
// always logs the given fields.
func WithFields(ctx context.Context, fields logrus.Fields) context.Context {
	return WithLogger(ctx, Logger(ctx).WithFields(fields))
}

// WithTime creates an entry from the standard logger and overrides the time of
// logs generated with it.
func WithTime(ctx context.Context, t time.Time) context.Context {
	return WithLogger(ctx, Logger(ctx).WithTime(t))
}

// WithContext creates an entry from the standard logger and adds a context to
// it.
func WithContext(ctx context.Context) context.Context {
	return WithLogger(ctx, Logger(ctx).WithContext(ctx))
}

// WithError creates an entry from the standard logger and adds an error to it,
// using the value defined in logrus.ErrorKey as key.
func WithError(ctx context.Context, err error) context.Context {
	return WithLogger(ctx, Logger(ctx).WithError(err))
}

// Debugf calls Logger(ctx).Debugf(format string, args...)
func Debugf(ctx context.Context, format string, args ...interface{}) {
	Logger(ctx).Debugf(format, args...)
}

// Infof calls Logger(ctx).Infof(format string, args...)
func Infof(ctx context.Context, format string, args ...interface{}) {
	Logger(ctx).Infof(format, args...)
}

// Printf calls Logger(ctx).Printf(format string, args...)
func Printf(ctx context.Context, format string, args ...interface{}) {
	Logger(ctx).Printf(format, args...)
}

// Warnf calls Logger(ctx).Warnf(format string, args...)
func Warnf(ctx context.Context, format string, args ...interface{}) {
	Logger(ctx).Warnf(format, args...)
}

// Warningf calls Logger(ctx).Warningf(format string, args...)
func Warningf(ctx context.Context, format string, args ...interface{}) {
	Logger(ctx).Warningf(format, args...)
}

// Errorf calls Logger(ctx).Errorf(format string, args...)
func Errorf(ctx context.Context, format string, args ...interface{}) {
	Logger(ctx).Errorf(format, args...)
}

// Fatalf calls Logger(ctx).Fatalf(format string, args...)
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	Logger(ctx).Fatalf(format, args...)
}

// Panicf calls Logger(ctx).Panicf(format string, args...)
func Panicf(ctx context.Context, format string, args ...interface{}) {
	Logger(ctx).Panicf(format, args...)
}

// Debug calls Logger(ctx).Debug(args...)
func Debug(ctx context.Context, args ...interface{}) {
	Logger(ctx).Debug(args...)
}

// Info calls Logger(ctx).Info(args...)
func Info(ctx context.Context, args ...interface{}) {
	Logger(ctx).Info(args...)
}

// Print calls Logger(ctx).Print(args...)
func Print(ctx context.Context, args ...interface{}) {
	Logger(ctx).Print(args...)
}

// Warn calls Logger(ctx).Warn(args...)
func Warn(ctx context.Context, args ...interface{}) {
	Logger(ctx).Warn(args...)
}

// Warning calls Logger(ctx).Warning(args...)
func Warning(ctx context.Context, args ...interface{}) {
	Logger(ctx).Warning(args...)
}

// Error calls Logger(ctx).Error(args...)
func Error(ctx context.Context, args ...interface{}) {
	Logger(ctx).Error(args...)
}

// Fatal calls Logger(ctx).Fatal(args...)
func Fatal(ctx context.Context, args ...interface{}) {
	Logger(ctx).Fatal(args...)
}

// Panic calls Logger(ctx).Panic(args...)
func Panic(ctx context.Context, args ...interface{}) {
	Logger(ctx).Panic(args...)
}

// Debugln calls Logger(ctx).Debugln(args...)
func Debugln(ctx context.Context, args ...interface{}) {
	Logger(ctx).Debugln(args...)
}

// Infoln calls Logger(ctx).Infoln(args...)
func Infoln(ctx context.Context, args ...interface{}) {
	Logger(ctx).Infoln(args...)
}

// Println calls Logger(ctx).Println(args...)
func Println(ctx context.Context, args ...interface{}) {
	Logger(ctx).Println(args...)
}

// Warnln calls Logger(ctx).Warnln(args...)
func Warnln(ctx context.Context, args ...interface{}) {
	Logger(ctx).Warnln(args...)
}

// Warningln calls Logger(ctx).Warningln(args...)
func Warningln(ctx context.Context, args ...interface{}) {
	Logger(ctx).Warningln(args...)
}

// Errorln calls Logger(ctx).Errorln(args...)
func Errorln(ctx context.Context, args ...interface{}) {
	Logger(ctx).Errorln(args...)
}

// Fatalln calls Logger(ctx).Fatalln(args...)
func Fatalln(ctx context.Context, args ...interface{}) {
	Logger(ctx).Fatalln(args...)
}

// Panicln calls Logger(ctx).Panicln(args...)
func Panicln(ctx context.Context, args ...interface{}) {
	Logger(ctx).Panicln(args...)
}
