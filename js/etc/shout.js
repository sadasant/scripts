// shout.js
// Shout out text files over your network

// Requires
var http = require('http')
  , fs   = require('fs')

// Simple Variables
var contentType = {'Content-Type':'text/plain'}
  , fileName    = process.argv[2]
  , fileContent

function readFile(req, res) {
  fs.readFile(fileName, 'utf-8', function(err, data) {
    if (err) {
      console.error("Could not open file %s", err)
      process.exit(1)
    }
    fileContent = data.toString('utf-8')
    return shout(req, res)
  })
}

function shout(req, res) {
  res.writeHead(200, contentType)
  res.write(fileContent)
  res.end()
}

// Server
http.createServer(readFile).listen(1337)
