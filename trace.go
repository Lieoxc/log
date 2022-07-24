package logs

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// WithContext sets ctx to log, for keeping tracing information.
func WithContext(ctx context.Context) Logger {
	return &traceLogger{
		ctx: ctx,
	}
}

type traceLogger struct {
	ctx context.Context
}

func (l *traceLogger) WithContext(ctx context.Context) Logger {
	if ctx == nil {
		return l
	}

	l.ctx = ctx
	return l
}

func (l *traceLogger) Info(format string, v ...interface{}) {

	userID, projectID, traceid, spanId := l.getContextInfo()
	content := fmt.Sprintf(format, v...)

	log.logWrite.Info(content, zap.Any("userID", userID),
		zap.Any("projectID", projectID),
		zap.Any("trace", traceid),
		zap.Any("span", spanId))

}

func (l *traceLogger) Error(format string, v ...interface{}) {

	userID, projectID, traceid, spanId := l.getContextInfo()

	content := fmt.Sprintf(format, v...)

	log.logWrite.Error(content, zap.Any("userID", userID),
		zap.Any("projectID", projectID),
		zap.Any("trace", traceid),
		zap.Any("span", spanId))

}

// func LogFatal(ctx context.Context, format string, v ...interface{}) {

// 	userID, projectID := getHeader(ctx)
// 	traceid := traceIdFromContext(ctx)
// 	spanId := spanIdFromContext(ctx)

// 	content := fmt.Sprintf(format, v...)

// 	log.logWrite.Fatal(content, zap.Any("userID", userID),
// 		zap.Any("projectID", projectID),
// 		zap.Any("trace", traceid),
// 		zap.Any("span", spanId))

// }
func (l *traceLogger) getContextInfo() (string, string, string, string) {
	userID, projectID := getHeader(l.ctx)
	traceid := traceIdFromContext(l.ctx)
	spanId := spanIdFromContext(l.ctx)
	return userID, projectID, traceid, spanId
}

func getHeader(ctx context.Context) (string, string) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ""
	}

	userID := headers.Get("userID")
	projectIDs := headers.Get("projectID")
	if len(projectIDs) > 0 && len(userID) > 0 {
		return userID[0], projectIDs[0]
	}
	return "", ""
}
func traceIdFromContext(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		return spanCtx.TraceID().String()
	}

	return ""
}
func spanIdFromContext(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasSpanID() {
		return spanCtx.SpanID().String()
	}

	return ""
}
