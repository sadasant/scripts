var http = require("http")
  , url  = require("url")
  , fs   = require("fs")


var pages = {
  "/"         : { "type" : "text/html"      , "file" : "test.html" }
, "/Shade.js" : { "type" : "text/javascript", "file" : "Shade.js"  }
}

var requestListener = function (req, res) {
  var path = url.parse(req.url).pathname
    , page = pages[path]
    , msg  = "404"
    , body = ""

  if (page === undefined) {

    // Get test
    if (path === "/get-test") {
      msg = "get test ok"
    }

    // Post test
    if (path === "/post-test") {
      req.on("data", function(data) {
        body += data
        // Flood?
        if (body.length > 1e6) {
          body = ""
          request.connection.destroy()
        }
      })
      req.on("end", function() {
        res.write(body === "test=true" ? "post test ok" : "")
        res.end()
      })
      return
    }

    // 404
    res.write(msg)
    res.end()
    return

  }

  // Plain Text
  res.writeHead(200, {"Content-Type": page.type})
  fs.readFile(page.file, "utf8", function(err, data) {
    res.write(data)
    res.end()
  })

}

var server = http.createServer(requestListener)
server.listen(8080)
