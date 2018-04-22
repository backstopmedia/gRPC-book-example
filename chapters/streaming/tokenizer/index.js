var fs = require('fs')

var async = require('async')
var grpc = require('grpc')

var messages = require('./proto/tokenizer_pb.js')
var services = require('./proto/tokenizer_grpc_pb.js')
var client = new services.TokenizerClient('localhost:8080', grpc.credentials.createInsecure())

function tokenizeFiles(files, done) {
  var call = client.tokenize()
  var error = null
  var wordsMap = {}

  call.on('data', function(resp) {
    resp.toObject().wordsMap.forEach(function(obj) {
      var word = obj[0]
      var count = obj[1]

      wordsMap[word] = (wordsMap[word] || 0) + count
    })
  })

  call.on('end', function() { done(error, wordsMap) })

  sendCalls = files.map(function(file) {
    return function(callback) {
      fs.readFile(file, function(err, data) {
        if (err != null) {
          callback(err, null)
          return
        }

        var req = new messages.TokenizeRequest()
        req.setFileContents(data.toString('base64'))

        call.write(req)
        callback(null, null)
      })
    }
  })

  async.parallel(sendCalls, function(err, _) {
    error = err
    call.end()
  })
}

tokenizeFiles(process.argv.slice(2), function(err, results) {
  if (err != null) {
    console.log('An Error occured:', err)
    process.exit(1)
  }

  console.log('All finished')
  console.log(results)
})
