#!/usr/bin/env ./node_modules/.bin/babel-node
/**
 * Get data from https://sfapi.co/ and normalize it
 * Definition diverges form this, so this is just for initial data
 */

import fetch from 'node-fetch'
import { writeFileSync } from 'fs'
import { generate as shortid } from 'shortid'

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

const rRemoveUrl = new RegExp('https://sfapi.co/api/(.+)/')

// this holds all cross-refed IDs
const ids = []

const main = async () => {
  const getNoun = async (noun, params = {}) => {
    let next = urls[noun] + handleParams(params)
    const results = []
    while (next !== null) {
      const r = await fetch(next).then(r => r.json())
      next = r.next
      const cleanedResults = r.results.map(result => {
        result.id = result.url.replace(rRemoveUrl, '$1').replace('/', '_')
        ids.push(result.id)
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

        // try my best to get data in a better format
        Object.keys(result).forEach(k => {
          // lots of misued strings
          if (typeof result[k] === 'string') {
            // edge-case fixes
            if (noun === 'people' && k === 'mass') {
              result[k] = parseFloat(result[k].replace(/([\d.]+)/g, '$1'))
              return
            }

            if (noun === 'planets' && k === 'gravity') {
              result[k] = parseFloat(result[k].replace(/([\d.]+)/g, '$1'))
              return
            }

            if (noun === 'vehicles' && k === 'consumables') {
              if (result[k] === 'Live food tanks') {
                result[k] = -1
                return
              }
              if (result[k] === '0' || result[k] === 'none') {
                result[k] = 0
                return
              }
              if (result[k] === 'unknown') {
                result[k] = undefined
                return
              }
              if (result[k].indexOf('month') !== -1) {
                result[k] = parseFloat(result[k].replace(/ months?/, '')) * 30
                return
              }
              if (result[k].indexOf('day') !== -1) {
                result[k] = parseFloat(result[k].replace(/ days?/, ''))
                return
              }
            }

            // the data uses strings with commas for numbers!
            if (result[k].indexOf(',') !== -1) {
              result[k] = result[k].replace(/(\d+),(?=\d{3}(\D|$))/g, '$1')
            }
            if (result[k] == parseInt(result[k])) {
              result[k] = parseInt(result[k])
              return
            }
            if (result[k] == parseFloat(result[k])) {
              result[k] = parseFloat(result[k])
              return
            }

            // undefined is modeled differently in different records
            if (result[k] === 'unknown') {
              result[k] = undefined
              return
            }
            if (result[k].toLowerCase() === 'n/a') {
              result[k] = undefined
            }
          }
        })
        return result
      })
      results.push.apply(results, cleanedResults)
    }
    return results
  }

  const urls = await fetch('https://sfapi.co/api/', {headers}).then(r => r.json())
  const out = {}
  await Promise.all(Object.keys(urls).map(async noun => {
    out[noun] = await getNoun(noun)
  }))

  // replace all cross-reffed IDs
  let outS = JSON.stringify(out, null, 2)
  ids.forEach(i => {
    outS = outS.replace(new RegExp(i, 'g'), shortid())
  })

  writeFileSync('../data.json', outS)
}

main()
