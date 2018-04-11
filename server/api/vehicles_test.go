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

type VehiclesTest struct {
	suite.Suite
	api *api.Server
	db  *mocks.Provider
}

func TestVehicles(t *testing.T) {
	suite.Run(t, new(VehiclesTest))
}

func (assert *VehiclesTest) SetupTest() {
	assert.db = new(mocks.Provider)
	assert.api = api.New(assert.db)
}

func (assert *VehiclesTest) TestGetVehicle() {
	assert.db.On("GetVehicleByID", mock.Anything, "someid").Return(&pb.Vehicle{Name: "Some Name"}, nil)

	req := &pb.GetVehicleRequest{Id: "someid"}
	resp, err := assert.api.GetVehicle(context.Background(), req)

	assert.NoError(err)
	assert.Equal("Some Name", resp.Vehicle.Name)
}

func (assert *VehiclesTest) TestGetVehicleError() {
	assert.db.On("GetVehicleByID", mock.Anything, "someid").Return(nil, errors.New("Boom"))

	req := &pb.GetVehicleRequest{Id: "someid"}
	resp, err := assert.api.GetVehicle(context.Background(), req)

	assert.Nil(resp)
	assert.EqualError(err, "Boom")
}

func (assert *VehiclesTest) TestListVehicles() {
	assert.db.On("GetVehicles", mock.Anything).Return(
		[]*pb.Vehicle{{Name: "Vehicle1"}, {Name: "Vehicle2"}},
		nil,
	)

	resp, err := assert.api.ListVehicles(context.Background(), new(pb.ListVehiclesRequest))

	assert.NoError(err)
	assert.Len(resp.Vehicles, 2)
	assert.Equal("Vehicle1", resp.Vehicles[0].Name)
	assert.Equal("Vehicle2", resp.Vehicles[1].Name)
}

func (assert *VehiclesTest) TestListVehiclesError() {
	assert.db.On("GetVehicles", mock.Anything).Return(nil, errors.New("Boom"))
	resp, err := assert.api.ListVehicles(context.Background(), new(pb.ListVehiclesRequest))

	assert.Nil(resp)
	assert.EqualError(err, "Boom")
}
