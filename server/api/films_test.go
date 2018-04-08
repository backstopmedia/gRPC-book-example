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

type FilmsTest struct {
	suite.Suite
	api *api.Server
	db  *mocks.Provider
}

func TestFilms(t *testing.T) {
	suite.Run(t, &FilmsTest{api: new(api.Server)})
}

func (assert *FilmsTest) SetupTest() {
	assert.db = new(mocks.Provider)
	assert.api = api.New(assert.db)
}

func (assert *FilmsTest) TestGetFilm() {
	assert.db.On("GetFilmByID", mock.Anything, "someid").Return(&pb.Film{Title: "Some Title"}, nil)

	req := &pb.GetFilmRequest{Id: "someid"}
	resp, err := assert.api.GetFilm(context.Background(), req)

	assert.NoError(err)
	assert.Equal("Some Title", resp.Film.Title)
}

func (assert *FilmsTest) TestGetFilmError() {
	assert.db.On("GetFilmByID", mock.Anything, "someid").Return(nil, errors.New("Boom"))

	req := &pb.GetFilmRequest{Id: "someid"}
	resp, err := assert.api.GetFilm(context.Background(), req)

	assert.Nil(resp)
	assert.EqualError(err, "Boom")
}

func (assert *FilmsTest) TestListFilms() {
	assert.db.On("GetFilms", mock.Anything).Return(
		[]*pb.Film{{Title: "Film1"}, {Title: "Film2"}},
		nil,
	)

	resp, err := assert.api.ListFilms(context.Background(), new(pb.ListFilmsRequest))

	assert.NoError(err)
	assert.Len(resp.Films, 2)
	assert.Equal("Film1", resp.Films[0].Title)
	assert.Equal("Film2", resp.Films[1].Title)
}

func (assert *FilmsTest) TestListFilmsError() {
	assert.db.On("GetFilms", mock.Anything).Return(nil, errors.New("Boom"))
	resp, err := assert.api.ListFilms(context.Background(), new(pb.ListFilmsRequest))

	assert.Nil(resp)
	assert.EqualError(err, "Boom")
}
