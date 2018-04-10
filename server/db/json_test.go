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

func (assert *JSONTest) TestGetPersonByID() {
	person, err := assert.db.GetPersonByID(context.Background(), "SyAbJp35ViM")
	assert.NoError(err)

	assert.Equal("Luke Skywalker", person.Name)
}

func (assert *JSONTest) TestGetPersonByIDNotFound() {
	person, err := assert.db.GetPersonByID(context.Background(), "NoGood")
	assert.Nil(person)
	assert.EqualError(err, "person with id 'NoGood' not found")
}

func (assert *JSONTest) TestGetPeople() {
	people, err := assert.db.GetPeople(context.Background())
	assert.NoError(err)

	assert.Len(people, 87)
}

func (assert *JSONTest) TestGetPlanetByID() {
	planet, err := assert.db.GetPlanetByID(context.Background(), "B17gkTncNoz")
	assert.NoError(err)

	assert.Equal("Alderaan", planet.Name)
}

func (assert *JSONTest) TestGetPlanetByIDNotFound() {
	planet, err := assert.db.GetPlanetByID(context.Background(), "NoGood")
	assert.Nil(planet)
	assert.EqualError(err, "planet with id 'NoGood' not found")
}

func (assert *JSONTest) TestGetPlanets() {
	planets, err := assert.db.GetPlanets(context.Background())
	assert.NoError(err)

	assert.Len(planets, 61)
}
