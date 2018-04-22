package api_test

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"bytes"
	"context"
	"errors"
	"os"
	"testing"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	"github.com/backstopmedia/gRPC-book-example/server/db"
)

func TestErrorsInterceptor(t *testing.T) {
	t.Run("Errors interceptor returns a codes.NotFound when a DB not found is returned", func(t *testing.T) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return nil, db.NotFoundErr{}
		}

		out, err := api.ErrorsInterceptor(nil, nil, nil, handler)
		assert.Nil(t, out)
		assert.Equal(t, codes.NotFound, grpc.Code(err))
	})

	t.Run("Errors interceptor returns the originl error if not applicable to intercept", func(t *testing.T) {
		expErr := errors.New("Boom")

		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return nil, expErr
		}

		out, err := api.ErrorsInterceptor(nil, nil, nil, handler)
		assert.Nil(t, out)
		assert.Equal(t, expErr, err)
	})
}

func TestLoggingInterceptor(t *testing.T) {
	buf := new(bytes.Buffer)
	logrus.SetOutput(buf)
	defer logrus.SetOutput(os.Stdout)

	// no-op handler
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	}

	info := &grpc.UnaryServerInfo{FullMethod: "boom.Boom/GetBoom"}
	_, err := api.LoggingInterceptor(nil, nil, info, handler)
	require.NoError(t, err, "error on logging interceptor")

	logLine := buf.String()
	assert.Contains(t, logLine, "method=boom.Boom/GetBoom")
	assert.Contains(t, logLine, "duration=")
	assert.Contains(t, logLine, "finished RPC")
}
