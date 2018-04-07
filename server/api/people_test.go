package api_test

import (
	"github.com/stretchr/testify/suite"

	"context"
	"testing"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
)

type PeopleTest struct {
	suite.Suite
	api *api.Server
}

func TestPeople(t *testing.T) {
	suite.Run(t, &PeopleTest{api: new(api.Server)})
}

func (assert *PeopleTest) TestGetPerson() {
	req := &pb.GetPersonRequest{Id: "someid"}
	resp, err := assert.api.GetPerson(context.Background(), req)

	assert.Nil(resp)
	assert.EqualError(err, "rpc error: code = Unimplemented desc = Not implemented yet")
}

func (assert *PeopleTest) TestListPeople() {
	resp, err := assert.api.ListPeople(context.Background(), new(pb.ListPeopleRequest))

	assert.Nil(resp)
	assert.EqualError(err, "rpc error: code = Unimplemented desc = Not implemented yet")
}
