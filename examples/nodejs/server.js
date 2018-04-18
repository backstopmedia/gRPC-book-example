const data = require('../../data.json')

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

const listHandler = (recordName, dataName) => {
  const records = ({[recordName]: data[dataName]})
  return records
}

// this could all be generated from a few nouns
// but here it is all spelled-out for illustration
module.exports = {
  swapi: {
    v1: {
      Starwars: {
        GetFilm: ({request: {id}}) => getHandler(id, 'film', 'films'),
        ListFilms: () => listHandler('films', 'films'),
        GetVehicle: ({request: {id}}) => getHandler(id, 'vehicle', 'vehicles'),
        ListVehicles: () => listHandler('vehicles', 'vehicles'),
        GetStarship: ({request: {id}}) => getHandler(id, 'starship', 'starships'),
        ListStarships: () => listHandler('starships', 'starships'),
        GetSpecies: ({request: {id}}) => getHandler(id, 'species', 'species'),
        ListSpecies: () => listHandler('species', 'species'),
        GetPlanet: ({request: {id}}) => getHandler(id, 'planet', 'planets'),
        ListPlanets: () => listHandler('planets', 'planets'),
        GetPerson: ({request: {id}}) => getHandler(id, 'person', 'people'),
        ListPeople: () => listHandler('people', 'people'),
        ListStarshipActions: () => {} // NO-OP
      }
    }
  }
}
