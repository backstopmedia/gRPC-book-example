const data = require('../data')
const _ = require('lodash')

/**
 * Simple demo resolvers implemented in node
 */

module.exports = {
  sfapi: {
    v1: {
      Starfriends: {
        GetFilm: ({ request: { id } }, cb) => {
          const results = _.find(data.films, {id})
          if (!results) {
            return cb({message: 'Record not found', code: 5})
          }
          cb(null, results)
        },

        ListFilms: (ctx, cb) => {
          const results = data.films
          cb(null, {results})
        },

        GetStarship: ({ request: { id } }, cb) => {
          const results = _.find(data.starships, {id})
          if (!results) {
            return cb({message: 'Record not found', code: 5})
          }
          cb(null, results)
        },

        ListStarships: (ctx, cb) => {
          const results = data.starships
          cb(null, {results})
        },

        GetVehicle: ({ request: { id } }, cb) => {
          const results = _.find(data.vehicles, {id})
          if (!results) {
            return cb({message: 'Record not found', code: 5})
          }
          cb(null, results)
        },

        ListVehicles: (ctx, cb) => {
          const results = data.vehicles
          cb(null, {results})
        },

        GetSpecies: ({ request: { id } }, cb) => {
          const results = _.find(data.species, {id})
          if (!results) {
            return cb({message: 'Record not found', code: 5})
          }
          cb(null, results)
        },

        ListSpecies: (ctx, cb) => {
          const results = data.species
          cb(null, {results})
        },

        GetPlanet: ({ request: { id } }, cb) => {
          const results = _.find(data.planets, {id})
          if (!results) {
            return cb({message: 'Record not found', code: 5})
          }
          cb(null, results)
        },

        ListPlanets: (ctx, cb) => {
          const results = data.planets
          cb(null, {results})
        },

        GetPerson: ({ request: { id } }, cb) => {
          const results = _.find(data.people, {id})
          if (!results) {
            return cb({message: 'Record not found', code: 5})
          }
          cb(null, results)
        },

        ListPeople: (ctx, cb) => {
          const results = data.people
          cb(null, {results})
        }
      }
    }
  }
}
