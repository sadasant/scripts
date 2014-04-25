var euler = require('./euler')

function solution(n) {
    var max = Math.pow(10, n) - 1
    var min = parseInt(9 * max / 10)
    for (var i = max;; i--) {
        for (var j = max; j > min; j--) {
            var p = ""+(i * j)
            if (p === p.split("").reverse("").join("")) {
                return p
            }
        }
    }
}

euler.Init(4, "Find the largest palindrome made from the product of two N-digit numbers.")
euler.PrintTime("N = 2 | Result: %d, Nanoseconds: %d", solution, 2)
euler.PrintTime("N = 3 | Result: %d, Nanoseconds: %d", solution, 3)
