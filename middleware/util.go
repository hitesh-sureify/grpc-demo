package middleware

import (
	"google.golang.org/grpc"
	"go.uber.org/zap"
)

func AddInterceptors(opts []grpc.ServerOption, uInterceptors []grpc.UnaryServerInterceptor, sInterceptors []grpc.StreamServerInterceptor) []grpc.ServerOption {
	opts = append(opts, grpc.ChainUnaryInterceptor(uInterceptors...))
	opts = append(opts, grpc.ChainStreamInterceptor(sInterceptors...))
	return opts
}

// Add grpc default middleware like logging and prometheus metrics
func GetGrpcMiddlewareOpts() []grpc.ServerOption {
	// gRPC server startup options
	opts := []grpc.ServerOption{}

	uInterceptors := []grpc.UnaryServerInterceptor{}
	sInterceptors := []grpc.StreamServerInterceptor{}

	// add middleware
	AddLogging(&zap.Logger, &uInterceptors, &sInterceptors)
	AddPrometheus(&uInterceptors, &sInterceptors)

	opts = AddInterceptors(opts, uInterceptors, sInterceptors)

	return opts
}
