var euler = module.exports

// var b_black   = "\x1b[30;1m"
// var b_red     = "\x1b[31;1m"
// var b_green   = "\x1b[32;1m"
// var b_blue    = "\x1b[34;1m"
// var b_magenta = "\x1b[35;1m"
// var b_cian    = "\x1b[36;1m"
// var b_white   = "\x1b[37;1m"
// var black     = "\x1b[30m"
// var red       = "\x1b[31m"
// var green     = "\x1b[32m"
// var blue      = "\x1b[34m"
// var magenta   = "\x1b[35m"
// var cian      = "\x1b[36m"
// var white     = "\x1b[37m"

var b_yellow = "\x1b[33;1m"
var yellow   = "\x1b[33m"
var normal   = "\x1b[0m"

euler.Init = function(n, text) {
    console.log("%sProject Euler, problem %d%s%s\n%s%s", b_yellow, n, normal, yellow, text, normal)
};

euler.PrintTime = function(msg, f) {
    var v = Array.prototype.slice.call(arguments, 2)
    var t = process.hrtime()
    var r = f.apply(undefined, v)
    t = process.hrtime(t)[1]
    console.log(msg, r, t)
};
