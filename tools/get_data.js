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

const main = async () => {
  const urls = await fetch('https://swapi.co/api/', {headers}).then(r => r.json())
  const getNoun = async (noun, params = {}) => {
    let next = urls[noun] + handleParams(params)
    const results = []
    while (next !== null) {
      const r = await fetch(next).then(r => r.json())
      next = r.next
      const cleanedResults = r.results.map(result => {
        result.id = parseInt(result.url.replace(new RegExp(`https://swapi.co/api/${noun}/([0-9]+)/`), '$1'))
        delete result.url

        Object.keys(nounMap).forEach(noun => {
          if (result[noun]) {
            if (Array.isArray(result[noun])) {
              result[noun] = result[noun].map(c => parseInt(c.replace(new RegExp(`https://swapi.co/api/${nounMap[noun]}/([0-9]+)/`), '$1')))
            } else {
              result[noun] = parseInt(result[noun].replace(new RegExp(`https://swapi.co/api/${nounMap[noun]}/([0-9]+)/`), '$1'))
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
