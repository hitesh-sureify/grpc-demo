package middleware

import (
	"context"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

const RequestIDKey string = "request-id"

// GetReqID returns a request ID from the given context if one is present.
// Returns the empty string if a request ID cannot be found.
func GetReqID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}

func GetContextRequestID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md[RequestIDKey]) > 0 {
		return md[RequestIDKey][0]
	}
	return GetReqID(ctx)
}

// codeToLevel redirects OK to DEBUG level logging instead of Info
func codeToLevel(code codes.Code) zapcore.Level {
	if code == codes.OK {
		// It is DEBUG
		return zap.DebugLevel
	}
	return grpc_zap.DefaultCodeToLevel(code)
}

func extractFields(fullMethod string, req interface{}) map[string]interface{} {
	//log.Printf("%v", req)
	return make(map[string]interface{})
}

func messageProducer(ctx context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field) {
	ctxzap.AddFields(ctx, zap.String(RequestIDKey, GetContextRequestID(ctx)))
	ctxzap.Extract(ctx).Check(level, msg).Write(
		zap.Error(err),
		zap.String("grpc.code", code.String()),
		duration,
	)
}


// AddLogging returns grpc.Server config option that turn on logging.
func AddLogging(logger *zap.Logger, uInterceptors *[]grpc.UnaryServerInterceptor, sInterceptors *[]grpc.StreamServerInterceptor) {
	// Shared options for the logger, with a custom gRPC code to log level function.
	o := []grpc_zap.Option{
		grpc_zap.WithLevels(codeToLevel),
		grpc_zap.WithMessageProducer(messageProducer),
	}

	// Make sure that log statements internal to gRPC library are logged using the zaplogger as well.
	grpc_zap.ReplaceGrpcLoggerV2(logger)

	*uInterceptors = append(*uInterceptors, grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(extractFields)))
	*uInterceptors = append(*uInterceptors, grpc_zap.UnaryServerInterceptor(logger, o...))

	*sInterceptors = append(*sInterceptors, grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)))
	*sInterceptors = append(*sInterceptors, grpc_zap.StreamServerInterceptor(logger, o...))

}