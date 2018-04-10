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

type PeopleTest struct {
	suite.Suite
	api *api.Server
	db  *mocks.Provider
}

func TestPeople(t *testing.T) {
	suite.Run(t, new(PeopleTest))
}

func (assert *PeopleTest) SetupTest() {
	assert.db = new(mocks.Provider)
	assert.api = api.New(assert.db)
}

func (assert *PeopleTest) TestGetPerson() {
	assert.db.On("GetPersonByID", mock.Anything, "someid").Return(&pb.Person{Name: "Some Name"}, nil)

	req := &pb.GetPersonRequest{Id: "someid"}
	resp, err := assert.api.GetPerson(context.Background(), req)

	assert.NoError(err)
	assert.Equal("Some Name", resp.Person.Name)
}

func (assert *PeopleTest) TestGetPersonError() {
	assert.db.On("GetPersonByID", mock.Anything, "someid").Return(nil, errors.New("Boom"))

	req := &pb.GetPersonRequest{Id: "someid"}
	resp, err := assert.api.GetPerson(context.Background(), req)

	assert.Nil(resp)
	assert.EqualError(err, "Boom")
}

func (assert *PeopleTest) TestListPeople() {
	assert.db.On("GetPeople", mock.Anything).Return(
		[]*pb.Person{{Name: "Name1"}, {Name: "Name2"}},
		nil,
	)

	resp, err := assert.api.ListPeople(context.Background(), new(pb.ListPeopleRequest))

	assert.NoError(err)
	assert.Len(resp.People, 2)
	assert.Equal("Name1", resp.People[0].Name)
	assert.Equal("Name2", resp.People[1].Name)
}

func (assert *PeopleTest) TestListPeopleError() {
	assert.db.On("GetPeople", mock.Anything).Return(nil, errors.New("Boom"))
	resp, err := assert.api.ListPeople(context.Background(), new(pb.ListPeopleRequest))

	assert.Nil(resp)
	assert.EqualError(err, "Boom")
}
