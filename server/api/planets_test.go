package api_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"context"
	"errors"
	"testing"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	"github.com/backstopmedia/gRPC-book-example/server/api/mocks"
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

type PlanetsTest struct {
	suite.Suite
	api *api.Server
	db  *mocks.Provider
}

func TestPlanets(t *testing.T) {
	suite.Run(t, new(PlanetsTest))
}

func (assert *PlanetsTest) SetupTest() {
	assert.db = new(mocks.Provider)
	assert.api = api.New(assert.db)
}

func (assert *PlanetsTest) TestGetPlanet() {
	assert.db.On("GetPlanetByID", mock.Anything, "someid").Return(&pb.Planet{Name: "Some Name"}, nil)

	req := &pb.GetPlanetRequest{Id: "someid"}
	resp, err := assert.api.GetPlanet(context.Background(), req)

	assert.NoError(err)
	assert.Equal("Some Name", resp.Planet.Name)
}

func (assert *PlanetsTest) TestGetPlanetError() {
	assert.db.On("GetPlanetByID", mock.Anything, "someid").Return(nil, errors.New("Boom"))

	req := &pb.GetPlanetRequest{Id: "someid"}
	resp, err := assert.api.GetPlanet(context.Background(), req)

	assert.Nil(resp)
	assert.EqualError(err, "Boom")
}

func (assert *PlanetsTest) TestListPlanets() {
	assert.db.On("GetPlanets", mock.Anything).Return(
		[]*pb.Planet{{Name: "Planet1"}, {Name: "Planet2"}},
		nil,
	)

	resp, err := assert.api.ListPlanets(context.Background(), new(pb.ListPlanetsRequest))

	assert.NoError(err)
	assert.Len(resp.Planets, 2)
	assert.Equal("Planet1", resp.Planets[0].Name)
	assert.Equal("Planet2", resp.Planets[1].Name)
}

func (assert *PlanetsTest) TestListPlanetsError() {
	assert.db.On("GetPlanets", mock.Anything).Return(nil, errors.New("Boom"))
	resp, err := assert.api.ListPlanets(context.Background(), new(pb.ListPlanetsRequest))

	assert.Nil(resp)
	assert.EqualError(err, "Boom")
}
