package db

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

const slowThreshold = 200 * time.Millisecond // same as GORM

// FieldLogger is the same as logrus.FieldLogger, but with WithContext that was
// accidentally omitted.
type FieldLogger interface {
	log.FieldLogger
	WithContext(ctx context.Context) *log.Entry
}

// Logger is an adapter from logrus' logger to GORM.
type Logger struct {
	Logger FieldLogger
}

func (l Logger) LogMode(_ gormlogger.LogLevel) gormlogger.Interface {
	return l // ignore log mode change requests from GORM :)
}

func (l Logger) Error(ctx context.Context, s string, i ...interface{}) {
	l.Logger.WithContext(ctx).Errorf(s, i...)
}

func (l Logger) Warn(ctx context.Context, s string, i ...interface{}) {
	l.Logger.WithContext(ctx).Warnf(s, i...)
}

func (l Logger) Info(ctx context.Context, s string, i ...interface{}) {
	l.Logger.WithContext(ctx).Infof(s, i...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	logger := l.Logger.WithContext(ctx).WithError(err)
	loggerFunc := logger.Infof

	switch {
	case err == gorm.ErrRecordNotFound:
		loggerFunc = logger.Warnf
	case err != nil:
		loggerFunc = logger.Errorf
	case elapsed > slowThreshold:
		logger = logger.WithError(fmt.Errorf("slow query >= %v", slowThreshold))
		loggerFunc = logger.Infof
	}

	sql, rows := fc()
	if rows == -1 {
		loggerFunc("[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, "-", sql)
	} else {
		loggerFunc("[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}
