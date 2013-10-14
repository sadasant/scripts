var connect = require('connect');
var path = __dirname;
connect.createServer(
    connect.static(path)
).listen(8080);
console.log("Serving "+path+" at: localhost:8080");
console.log("Run at: http://localhost:8080/index.html");
