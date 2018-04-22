const emitonoff = require('emitonoff')
const data = require('../../data.json')

// I could read the proto directly for the enum values, but this is faster/easier
const StarshipActions = [
  'TOOKOFF',
  'LANDED',
  'HYPERDRIVE',
  'HIDING_IN_A_MOUTH'
]

// emit events on this to inform listeners
const starshipTracker = emitonoff()

// fake ship events are happening!
setInterval(() => {
  const action = StarshipActions[ (StarshipActions.length * Math.random()) | 0 ]
  const starship = data.starships[ (data.starships.length * Math.random()) | 0 ]
  starshipTracker.emit('StarshipAction', {action, starship})
}, 1000)

const getHandler = (id, recordName, dataName) => {
  const record = data[dataName].find(r => r.id === id)
  if (record) {
    return {[recordName]: record}
  } else {
    const err = new Error('Not found.')
    err.code = 5
    return Promise.reject(err)
  }
}

const listHandler = (recordName) => ({[recordName]: data[recordName]})

module.exports = {
  sfapi: {
    v1: {
      Starfriends: {
        GetFilm: ({request: {id}}) => getHandler(id, 'film', 'films'),
        ListFilms: () => listHandler('films'),
        GetVehicle: ({request: {id}}) => getHandler(id, 'vehicle', 'vehicles'),
        ListVehicles: () => listHandler('vehicles'),
        GetStarship: ({request: {id}}) => getHandler(id, 'starship', 'starships'),
        ListStarships: () => listHandler('starships'),
        GetSpecies: ({request: {id}}) => getHandler(id, 'species', 'species'),
        ListSpecies: () => listHandler('species'),
        GetPlanet: ({request: {id}}) => getHandler(id, 'planet', 'planets'),
        ListPlanets: () => listHandler('planets'),
        GetPerson: ({request: {id}}) => getHandler(id, 'person', 'people'),
        ListPeople: () => listHandler('people'),
        ListStarshipActions: (call) => {
          const handler = ev => call.write(ev)
          starshipTracker.on('StarshipAction', handler)
          call.on('cancelled', () => {
            starshipTracker.off('StarshipAction', handler)
            call.end()
          })
        }
      }
    }
  }
}
