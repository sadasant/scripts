var euler = require('./euler')

function solution(n) {
    var m = 3, l
    while (Math.sqrt(m) < n) {
        if (n%m === 0) {
            n = n/m
            l = m
        }
        m += 2
    }
    return l
}

euler.Init(3, "What is the largest prime factor of the number 600851475143 ?")
euler.PrintTime("Result: %d, Nanoseconds: %d\n", solution, 600851475143)
