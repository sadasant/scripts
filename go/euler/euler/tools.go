package euler

import "io/ioutil"
import "net/http"

// Sieve of Erastosthenes
// Ignoring even numbers.
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

// Download files with GET
func Download(path string) string {
	Printf("Downloading %v\n", path)
	resp, err := http.Get(path);
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	Printf("Downloaded %v\n", path)
	return string(body[:])
}
