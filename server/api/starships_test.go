package api_test

import (
	"github.com/stretchr/testify/suite"

	"context"
	"testing"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

type StarshipsTest struct {
	suite.Suite
	api *api.Server
}

func TestStarships(t *testing.T) {
	suite.Run(t, &StarshipsTest{api: new(api.Server)})
}

func (assert *StarshipsTest) TestGetStarship() {
	req := &pb.GetStarshipRequest{Id: "someid"}
	resp, err := assert.api.GetStarship(context.Background(), req)

	assert.Nil(resp)
	assert.EqualError(err, "rpc error: code = Unimplemented desc = Not implemented yet")
}

func (assert *StarshipsTest) TestListStarships() {
	resp, err := assert.api.ListStarships(context.Background(), new(pb.ListStarshipsRequest))

	assert.Nil(resp)
	assert.EqualError(err, "rpc error: code = Unimplemented desc = Not implemented yet")
}
