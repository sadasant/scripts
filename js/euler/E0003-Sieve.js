// Sieve of Erastosthenes
// Ignoring even numbers.
function primesUpTo(n) {
  var notprime = []
  var primes = []
  var sqrt = Math.sqrt(n)
  for (var i = 3; i <= sqrt; i += 2) {
    if (!notprime[i]) {
      primes.push(i)
      for (var j = i; j <= sqrt; j += i) {
        notprime[j] = true
      }
    }
  }
  return primes
}

function largestPrimeDivisorOf(n) {
  var primes = primesUpTo(Math.ceil(n/3)+1)
  for (var i = primes.length - 1; i > 0; i--) {
    if (n%primes[i] === 0) return primes[i]
  }
}

console.log(largestPrimeDivisorOf(600851475143))
