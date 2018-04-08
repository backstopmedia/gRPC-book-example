package db_test

import (
	"github.com/stretchr/testify/suite"

	"context"
	"io/ioutil"
	"testing"

	"github.com/backstopmedia/gRPC-book-example/server/db"
)

type JSONTest struct {
	suite.Suite
	db db.Provider
}

func TestJSON(t *testing.T) {
	suite.Run(t, new(JSONTest))
}

func (assert *JSONTest) SetupSuite() {
	bytes, err := ioutil.ReadFile("../../data.json")
	assert.NoError(err)

	assert.db, err = db.NewJSONProvider(bytes)
	assert.NoError(err)
}

func (assert *JSONTest) TestInitializationError() {
	db, err := db.NewJSONProvider([]byte(`{"films": {}}`))
	assert.Nil(db)
	assert.Error(err)
}

func (assert *JSONTest) TestGetFilmByID() {
	film, err := assert.db.GetFilmByID(context.Background(), "S1vZy63c4oz")
	assert.NoError(err)

	assert.Equal("A New Hope", film.Title)
}

func (assert *JSONTest) TestGetFilmByIDNotFound() {
	film, err := assert.db.GetFilmByID(context.Background(), "ThisDoesNotExist")
	assert.Nil(film)
	assert.EqualError(err, "film with id 'ThisDoesNotExist' not found")
}

func (assert *JSONTest) TestGetFilms() {
	films, err := assert.db.GetFilms(context.Background())
	assert.NoError(err)

	assert.Len(films, 7)
}
