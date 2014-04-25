var euler = require('./euler')

function solution(n) {
    var a = 0, b = 1, s = 0, c
    while (b < n) {
        c = b
        b = a+b
        a = c
        if (b%2 === 0) s += b
    }
    return s;
}

euler.Init(2, "By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.")
euler.PrintTime("Result: %d, Nanoseconds: %d", solution, 4e6)
