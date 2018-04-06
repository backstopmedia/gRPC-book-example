/**
 * Get data from https://swapi.co/
 */

import fetch from 'node-fetch'
import { writeFileSync } from 'fs'

const headers = {
  Accept: 'application/json'
}

// make a get query out of an object
const handleParams = params => {
  const o = Object.keys(params)
  if (o.length) {
    const p = o.map(k => encodeURIComponent(k) + '=' + encodeURIComponent(params[k]))
    return `?${p.join('&')}`
  } else {
    return ''
  }
}

// these are the nouns as they are seen in other nouns
const nounMap = {
  'characters': 'people',
  'residents': 'people',
  'pilots': 'people',
  'homeworld': 'planets',
  'people': 'people',
  'planets': 'planets',
  'starships': 'starships',
  'vehicles': 'vehicles',
  'species': 'species',
  'films': 'films'
}

const rRemoveUrl = new RegExp('https://swapi.co/api/(.+)/')

const main = async () => {
  const urls = await fetch('https://swapi.co/api/', {headers}).then(r => r.json())
  const getNoun = async (noun, params = {}) => {
    let next = urls[noun] + handleParams(params)
    const results = []
    while (next !== null) {
      const r = await fetch(next).then(r => r.json())
      next = r.next
      const cleanedResults = r.results.map(result => {
        result.id = result.url.replace(rRemoveUrl, '$1').replace('/', '_')
        delete result.url
        delete result.created
        delete result.edited

        Object.keys(nounMap).forEach(noun => {
          if (result[noun]) {
            if (Array.isArray(result[noun])) {
              result[noun] = result[noun].map(c => c.replace(rRemoveUrl, '$1').replace('/', '_'))
            } else {
              result[noun] = result[noun].replace(rRemoveUrl, '$1').replace('/', '_')
            }
          }
        })
        return result
      })
      results.push.apply(results, cleanedResults)
    }
    return results
  }
  const out = {}
  await Promise.all(Object.keys(urls).map(async noun => {
    out[noun] = await getNoun(noun)
  }))
  writeFileSync('../data.json', JSON.stringify(out, null, 2))
}

main()
