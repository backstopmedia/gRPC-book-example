package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	"github.com/backstopmedia/gRPC-book-example/server/api/mocks"
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

type StarshipActionsTest struct {
	suite.Suite
	api    *api.Server
	db     *mocks.Provider
	stream *mocks.Starfriends_ListStarshipActionsServer
}

func TestStarshipActions(t *testing.T) {
	suite.Run(t, new(StarshipActionsTest))
}

func (assert *StarshipActionsTest) SetupTest() {
	assert.db = new(mocks.Provider)
	assert.stream = new(mocks.Starfriends_ListStarshipActionsServer)
	assert.api = api.New(assert.db)
}

func (assert *StarshipActionsTest) TestListStarshipActions() {
	assert.stream.On("Context").Return(context.Background())
	starship := &pb.Starship{}
	assert.db.On("GetStarshipByID", mock.Anything, mock.Anything).Return(starship, nil)

	req := &pb.ListStarshipActionsRequest{}

	actions := []pb.StarshipAction_Action{
		pb.StarshipAction_TOOKOFF,
		pb.StarshipAction_HYPERDRIVE,
		pb.StarshipAction_LANDED,
		pb.StarshipAction_HIDING_IN_A_MOUTH,
	}
	for _, action := range actions {
		action := &pb.StarshipAction{
			Action:   action,
			Starship: starship,
		}
		assert.stream.On("Send", action).Return(nil)
	}

	err := assert.api.ListStarshipActions(req, assert.stream)
	assert.NoError(err)
	assert.stream.AssertExpectations(assert.T())
}
