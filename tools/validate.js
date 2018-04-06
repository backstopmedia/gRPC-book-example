/**
 * This will tell you if there are problems with your data/idl
 */
import protobuf from 'protobufjs'
import { join } from 'path'
import { singularize, classify } from 'inflection'
import data from '../data.json'

const p = protobuf.loadSync(join(__dirname, '..', 'proto', 'swapi.proto'))

Object.keys(data)
  .forEach(n => {
    const Type = p.lookupType(classify(singularize(n)))
    data[n].forEach(record => {
      const v = Type.verify(record)
      if (v) {
        const f = v.split(': ')[0]
        console.log(n, record.id, f, v, 'got', record[f])
      }
    })
  })
