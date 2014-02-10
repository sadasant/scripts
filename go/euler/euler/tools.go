package euler

// Sieve of Erastosthenes
func PrimesUpTo(n int) []int {
	bools := make([]bool, n)
	primes := make([]int, n)
	primes[0] = 2
	c := 1
	for i := 3; i < n; i += 2 {
		if !bools[i] {
			primes[c] = i
			c++
			for j := i; j < n; j += i {
				bools[j] = true
			}
		}
	}
	return primes[:c]
}
