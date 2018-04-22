package api_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
	"errors"
	"testing"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	"github.com/backstopmedia/gRPC-book-example/server/api/mocks"
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

type SpeciesTest struct {
	suite.Suite
	api *api.Server
	db  *mocks.Provider
}

func TestSpecies(t *testing.T) {
	suite.Run(t, new(SpeciesTest))
}

func (assert *SpeciesTest) SetupTest() {
	assert.db = new(mocks.Provider)
	assert.api = api.New(assert.db)
}

func (assert *SpeciesTest) TestGetSpecies() {
	assert.db.On("GetSpeciesByID", mock.Anything, "someid").Return(&pb.Species{Name: "Some Name"}, nil)

	req := &pb.GetSpeciesRequest{Id: "someid"}
	resp, err := assert.api.GetSpecies(context.Background(), req)

	assert.NoError(err)
	assert.Equal("Some Name", resp.Species.Name)
}

func (assert *SpeciesTest) TestGetSpeciesError() {
	assert.db.On("GetSpeciesByID", mock.Anything, "someid").Return(nil, errors.New("Boom"))

	req := &pb.GetSpeciesRequest{Id: "someid"}
	resp, err := assert.api.GetSpecies(context.Background(), req)

	assert.Nil(resp)
	assert.EqualError(err, "Boom")
}

func (assert *SpeciesTest) TestListSpecies() {
	assert.db.On("GetSpecies", mock.Anything).Return(
		[]*pb.Species{{Name: "Species1"}, {Name: "Species2"}},
		nil,
	)

	resp, err := assert.api.ListSpecies(context.Background(), new(pb.ListSpeciesRequest))

	assert.NoError(err)
	assert.Len(resp.Species, 2)
	assert.Equal("Species1", resp.Species[0].Name)
	assert.Equal("Species2", resp.Species[1].Name)
}

func (assert *SpeciesTest) TestListSpeciesError() {
	assert.db.On("GetSpecies", mock.Anything).Return(nil, errors.New("Boom"))
	resp, err := assert.api.ListSpecies(context.Background(), new(pb.ListSpeciesRequest))

	assert.Nil(resp)
	assert.EqualError(err, "Boom")
}

func (assert *SpeciesTest) TestValidateSpeciesError() {
	resp, err := assert.api.ValidateSpecies(context.Background(), new(pb.ValidateSpeciesRequest))

	assert.Nil(resp)
	assert.Equal(codes.InvalidArgument, status.Code(err))

	status, ok := status.FromError(err)
	assert.True(ok)

	assert.Require().Len(status.Details(), 1)
	d := status.Details()[0]

	invalidKey, ok := d.(*pb.InvalidKey)
	assert.True(ok)
	assert.Equal(invalidKey.GetKey(), "name")
	assert.Equal(invalidKey.GetMessage(), "The name provided is not long enough")
}
