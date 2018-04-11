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

type StarshipsTest struct {
	suite.Suite
	api *api.Server
	db  *mocks.Provider
}

func TestStarships(t *testing.T) {
	suite.Run(t, new(StarshipsTest))
}

func (assert *StarshipsTest) SetupTest() {
	assert.db = new(mocks.Provider)
	assert.api = api.New(assert.db)
}

func (assert *StarshipsTest) TestGetStarship() {
	assert.db.On("GetStarshipByID", mock.Anything, "someid").Return(&pb.Starship{Name: "Some Name"}, nil)

	req := &pb.GetStarshipRequest{Id: "someid"}
	resp, err := assert.api.GetStarship(context.Background(), req)

	assert.NoError(err)
	assert.Equal("Some Name", resp.Starship.Name)
}

func (assert *StarshipsTest) TestGetStarshipError() {
	assert.db.On("GetStarshipByID", mock.Anything, "someid").Return(nil, errors.New("Boom"))

	req := &pb.GetStarshipRequest{Id: "someid"}
	resp, err := assert.api.GetStarship(context.Background(), req)

	assert.Nil(resp)
	assert.EqualError(err, "Boom")
}

func (assert *StarshipsTest) TestListStarships() {
	assert.db.On("GetStarships", mock.Anything).Return(
		[]*pb.Starship{{Name: "Starship1"}, {Name: "Starship2"}},
		nil,
	)

	resp, err := assert.api.ListStarships(context.Background(), new(pb.ListStarshipsRequest))

	assert.NoError(err)
	assert.Len(resp.Starships, 2)
	assert.Equal("Starship1", resp.Starships[0].Name)
	assert.Equal("Starship2", resp.Starships[1].Name)
}

func (assert *StarshipsTest) TestListStarshipsError() {
	assert.db.On("GetStarships", mock.Anything).Return(nil, errors.New("Boom"))
	resp, err := assert.api.ListStarships(context.Background(), new(pb.ListStarshipsRequest))

	assert.Nil(resp)
	assert.EqualError(err, "Boom")
}
