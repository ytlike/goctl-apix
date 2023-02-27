package graylog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

type GrayLogWriter struct {
	logger *logrus.Logger
}

func NewGrayLogWriter(opts ...func(logger *logrus.Logger)) logx.Writer {
	host := os.Getenv("graylog.host")
	port := os.Getenv("graylog.port")
	namespace := os.Getenv("namespace")
	//appName := serverCenter.ServerName
	appName := "taskCenter-taskapi"

	logger := logrus.New()
	hook := NewGraylogHook(fmt.Sprintf("%s:%s", host, port), map[string]interface{}{"app_name": appName, "namespace": namespace})
	logger.Hooks.Add(hook)

	for _, opt := range opts {
		opt(logger)
	}

	return &GrayLogWriter{
		logger: logger,
	}
}

func (w *GrayLogWriter) Alert(v interface{}) {
	w.logger.Error(v)
}

func (w *GrayLogWriter) Close() error {
	w.logger.Exit(0)
	return nil
}

func (w *GrayLogWriter) Error(v interface{}, fields ...logx.LogField) {
	w.logger.WithFields(toLogrusFields(fields...)).Error(v)
}

func (w *GrayLogWriter) Info(v interface{}, fields ...logx.LogField) {
	w.logger.WithFields(toLogrusFields(fields...)).Info(v)
}

func (w *GrayLogWriter) Severe(v interface{}) {
	w.logger.Fatal(v)
}

func (w *GrayLogWriter) Slow(v interface{}, fields ...logx.LogField) {
	w.logger.WithFields(toLogrusFields(fields...)).Warn(v)
}

func (w *GrayLogWriter) Stack(v interface{}) {
	w.logger.Error(v)
}

func (w *GrayLogWriter) Stat(v interface{}, fields ...logx.LogField) {
	w.logger.WithFields(toLogrusFields(fields...)).Info(v)
}

func toLogrusFields(fields ...logx.LogField) logrus.Fields {
	logrusFields := make(logrus.Fields)
	for _, field := range fields {
		if field.Key == "trace" {
			logrusFields["traceId"] = field.Value
		} else if field.Key == "span" {
			logrusFields["spanId"] = field.Value
		} else {
			logrusFields[field.Key] = field.Value
		}
	}
	return logrusFields
}
