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
	film, err := assert.db.GetFilmByID(context.Background(), "NoGood")
	assert.Nil(film)
	assert.EqualError(err, "film with id 'NoGood' not found")
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

func (assert *JSONTest) TestGetSpeciesByID() {
	species, err := assert.db.GetSpeciesByID(context.Background(), "HJFy629NoG")
	assert.NoError(err)

	assert.Equal("Hutt", species.Name)
}

func (assert *JSONTest) TestGetSpeciesByIDNotFound() {
	species, err := assert.db.GetSpeciesByID(context.Background(), "NoGood")
	assert.Nil(species)
	assert.EqualError(err, "species with id 'NoGood' not found")
}

func (assert *JSONTest) TestGetSpecies() {
	species, err := assert.db.GetSpecies(context.Background())
	assert.NoError(err)

	assert.Len(species, 37)
}

func (assert *JSONTest) TestGetStarshipByID() {
	starship, err := assert.db.GetStarshipByID(context.Background(), "BkpeJa25Voz")
	assert.NoError(err)

	assert.Equal("Executor", starship.Name)
}

func (assert *JSONTest) TestGetStarshipByIDNotFound() {
	starship, err := assert.db.GetStarshipByID(context.Background(), "NoGood")
	assert.Nil(starship)
	assert.EqualError(err, "starship with id 'NoGood' not found")
}

func (assert *JSONTest) TestGetStarships() {
	starships, err := assert.db.GetStarships(context.Background())
	assert.NoError(err)

	assert.Len(starships, 37)
}

func (assert *JSONTest) TestGetVehicleByID() {
	vehicle, err := assert.db.GetVehicleByID(context.Background(), "Skka2cNjf")
	assert.NoError(err)

	assert.Equal("Sand Crawler", vehicle.Name)
}

func (assert *JSONTest) TestGetVehicleByIDNotFound() {
	vehicle, err := assert.db.GetVehicleByID(context.Background(), "NoGood")
	assert.Nil(vehicle)
	assert.EqualError(err, "vehicle with id 'NoGood' not found")
}

func (assert *JSONTest) TestGetVehicles() {
	vehicles, err := assert.db.GetVehicles(context.Background())
	assert.NoError(err)

	assert.Len(vehicles, 39)
}
