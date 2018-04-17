const data = require('../../data.json')
const { find } = require('lodash')

const notFound = () => {
  const err = new Error('Not found.')
  err.code = 5
  return Promise.reject(err)
}

// this could all be generated from a few nouns
// but here it is all spelled-out for illustration
module.exports = {
  swapi: {
    v1: {
      Starwars: {
        GetFilm: ({request: {id}}) => {
          const film = find(data.films, {id})
          if (film) {
            return { film }
          } else {
            notFound()
          }
        },
        ListFilms: () => data.films,
        GetVehicle: ({request: {id}}) => {
          const vehicle = find(data.vehicles, {id})
          if (vehicle) {
            return { vehicle }
          } else {
            notFound()
          }
        },
        ListVehicles: () => data.vehicles,
        GetStarship: ({request: {id}}) => {
          const starship = find(data.starships, {id})
          if (starship) {
            return { starship }
          } else {
            notFound()
          }
        },
        ListStarships: () => data.starships,
        GetSpecies: ({request: {id}}) => {
          const species = find(data.species, {id})
          if (species) {
            return { species }
          } else {
            notFound()
          }
        },
        ListSpecies: () => data.species,
        GetPlanet: ({request: {id}}) => {
          const planet = find(data.planets, {id})
          if (planet) {
            return { planet }
          } else {
            notFound()
          }
        },
        ListPlanets: () => data.planets,
        GetPerson: ({request: {id}}) => {
          const person = find(data.people, {id})
          if (person) {
            return { person }
          } else {
            notFound()
          }
        },
        ListPeople: () => data.people,
        ListStarshipActions: () => {} // NO-OP
      }
    }
  }
}
