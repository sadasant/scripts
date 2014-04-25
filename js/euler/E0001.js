var euler = require('./euler');

function solution(n) {
    var s = 0;
    for (var i = 0; i < n; i++) {
        if (i%3*i%5 === 0) s += i;
    }
    return s;
}

euler.Init(1, "Find the sum of all the multiples of 3 or 5 below 1000.");
euler.PrintTime("Result: %d, Nanoseconds: %d\n", solution, 1000);
