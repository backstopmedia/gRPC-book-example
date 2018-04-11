package api_test

import (
	"context"
	"errors"
	"testing"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func TestErrorsInterceptor(t *testing.T) {
	t.Run("Errros interceptor returns a codes.NotFound when a DB not found is returned", func(t *testing.T) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return nil, api.NotFoundErr{}
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
