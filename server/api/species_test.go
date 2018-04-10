package api_test

import (
	"github.com/stretchr/testify/suite"

	"context"
	"testing"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

type SpeciesTest struct {
	suite.Suite
	api *api.Server
}

func TestSpecies(t *testing.T) {
	suite.Run(t, &SpeciesTest{api: new(api.Server)})
}

func (assert *SpeciesTest) TestGetSpecies() {
	req := &pb.GetSpeciesRequest{Id: "someid"}
	resp, err := assert.api.GetSpecies(context.Background(), req)

	assert.Nil(resp)
	assert.EqualError(err, "rpc error: code = Unimplemented desc = Not implemented yet")
}

func (assert *SpeciesTest) TestListSpecies() {
	resp, err := assert.api.ListSpecies(context.Background(), new(pb.ListSpeciesRequest))

	assert.Nil(resp)
	assert.EqualError(err, "rpc error: code = Unimplemented desc = Not implemented yet")
}
