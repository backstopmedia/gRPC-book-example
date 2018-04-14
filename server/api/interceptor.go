package api

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"context"
	"time"

	"github.com/backstopmedia/gRPC-book-example/server/db"
)

// Interceptors implements the grpc.UnaryServerInteceptor function to add
// interceptors around all gRPC calls
func Interceptors() grpc.UnaryServerInterceptor {
	return grpc_middleware.ChainUnaryServer(
		LoggingInterceptor,
		ErrorsInterceptor,
	)
}

// ErrorsInterceptor adds error type checking to see if there are any known types
// what we return different grpc error codes for, for example: NotFound resources.
func ErrorsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {
	out, err = handler(ctx, req)

	switch tErr := err.(type) {
	case db.NotFoundErr:
		return out, grpc.Errorf(codes.NotFound, tErr.Error())
	}

	return out, err
}

// LoggingInterceptor adds logging around every gRPC call. It includes the method name and timing information.
// if the given handler raises an error, it also appends that to a key.
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {
	entry := logrus.WithField("method", info.FullMethod)
	start := time.Now()
	out, err = handler(ctx, req)
	duration := time.Since(start)

	if err != nil {
		entry = entry.WithError(err)
	}

	entry.WithField("duration", duration.String()).Info("finished RPC")

	return out, err
}
