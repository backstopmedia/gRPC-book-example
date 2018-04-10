package api_test

import (
	"github.com/stretchr/testify/suite"

	"context"
	"testing"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

type PlanetsTest struct {
	suite.Suite
	api *api.Server
}

func TestPlanets(t *testing.T) {
	suite.Run(t, &PlanetsTest{api: new(api.Server)})
}

func (assert *PlanetsTest) TestGetPlanet() {
	req := &pb.GetPlanetRequest{Id: "someid"}
	resp, err := assert.api.GetPlanet(context.Background(), req)

	assert.Nil(resp)
	assert.EqualError(err, "rpc error: code = Unimplemented desc = Not implemented yet")
}

func (assert *PlanetsTest) TestListPlanets() {
	resp, err := assert.api.ListPlanets(context.Background(), new(pb.ListPlanetsRequest))

	assert.Nil(resp)
	assert.EqualError(err, "rpc error: code = Unimplemented desc = Not implemented yet")
}
