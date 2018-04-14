#!/usr/bin/env ./node_modules/.bin/babel-node
/**
 * This will tell you if there are problems with your data/idl
 */
import protobuf from 'protobufjs'
import { join } from 'path'
import { singularize, classify } from 'inflection'
import chalk from 'chalk'
import data from '../data.json'

const p = protobuf.loadSync(join(__dirname, '..', 'proto', 'sfapi.proto'))

Object.keys(data)
  .forEach(n => {
    const Type = p.lookupType(classify(singularize(n)))
    data[n].forEach(record => {
      const v = Type.verify(record)
      const id = `${record.id.padEnd(12, ' ')} : ${Type.name.padEnd(8, ' ')}`
      if (v) {
        const f = v.split(': ')[0]
        console.log(chalk.red(`${id}: ${v} got ${JSON.stringify(record[f])}`))
      } else {
        console.log(chalk.green(`${id}: OK`))
      }
    })
  })
