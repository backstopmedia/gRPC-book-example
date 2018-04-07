package api_test

import (
	"github.com/stretchr/testify/suite"

	"context"
	"testing"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

type FilmsTest struct {
	suite.Suite
	api *api.Server
}

func TestFilms(t *testing.T) {
	suite.Run(t, &FilmsTest{api: new(api.Server)})
}

func (assert *FilmsTest) TestGetFilm() {
	req := &pb.GetFilmRequest{Id: "someid"}
	resp, err := assert.api.GetFilm(context.Background(), req)

	assert.Nil(resp)
	assert.EqualError(err, "rpc error: code = Unimplemented desc = Not implemented yet")
}

func (assert *FilmsTest) TestListFilms() {
	resp, err := assert.api.ListFilms(context.Background(), new(pb.ListFilmsRequest))

	assert.Nil(resp)
	assert.EqualError(err, "rpc error: code = Unimplemented desc = Not implemented yet")
}
