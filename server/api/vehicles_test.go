package api_test

import (
	"github.com/stretchr/testify/suite"

	"context"
	"testing"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

type VehiclesTest struct {
	suite.Suite
	api *api.Server
}

func TestVehicles(t *testing.T) {
	suite.Run(t, &VehiclesTest{api: new(api.Server)})
}

func (assert *VehiclesTest) TestGetVehicle() {
	req := &pb.GetVehicleRequest{Id: "someid"}
	resp, err := assert.api.GetVehicle(context.Background(), req)

	assert.Nil(resp)
	assert.EqualError(err, "rpc error: code = Unimplemented desc = Not implemented yet")
}

func (assert *VehiclesTest) TestListVehicles() {
	resp, err := assert.api.ListVehicles(context.Background(), new(pb.ListVehiclesRequest))

	assert.Nil(resp)
	assert.EqualError(err, "rpc error: code = Unimplemented desc = Not implemented yet")
}
