package shortcode

import (
	"math/rand"
	"strings"
	"time"
)

var vowels string = "aeiou"
var consonants string = "bcdfghjklmnpqrstvwxyz"

func Generate(length int) (code string) {
	var r int32
	var letter string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		switch i % 2 {
		case 0:
			r = rand.Int31n(int32(len(consonants)))
			letter = string(consonants[r])
		default:
			r = rand.Int31n(int32(len(vowels)))
			letter = string(vowels[r])
		}
		if r = rand.Int31n(2); r == 0 {
			letter = strings.ToUpper(letter)
		}
		code += letter
	}
	return code
}
