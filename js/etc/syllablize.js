// syllablize.js
//
// The idea is to build legible URLs by mixing
// single letter syllables, double character syllables
// and triple character syllables only with lower
// case letters, currently starts with a matrix of
// 1765 elements and works as converting decimals to hex.
//
// Ussage:
//
//     node syllablize.js 1000
//     Gef
//
// By Daniel Rodríguez
// <http://sadasant.com/>

var m          = []
  , vocals     = 'aeiou'
  , consonants = 'bcdfghjklmnpqrstvwxyz'
  , voc, cons, cons2
  , i, j, k
  , base

i = j = k = 0

// Generating the Matrix
// =====================

// Adding all the lower and upper vocals.
m = vocals.split('')

// Mixing them with two letter syllabes
// composed by [consonant][vocal] and [vocal][consonant]
// and later, adding up the three letter syllabes
// made up with [consonant][vocal][consonant].
while (voc = vocals[i+=1]) {
  j = 0
  while (cons = consonants[j+=1]) {
    m[m.length] = cons + voc
    m[m.length] = voc + cons
    k = 0
    while (cons2 = consonants[k+=1]) {
      m[m.length] = cons + voc + cons2
    }
  }
}

// Storing the current length (1765)
base = m.length

// Setting up the syllablize function
function syllablize(n) {
  var s = ''
    , r
  n >>= 0
  if (!n) {
    return 0
  }
  while (n > 0) {
    r = n % base
    n = (n / base) >> 0
    s = m[r] + s
  }
  return s
}

// Allowing manual tests from the console
// and exporting it for module ussage.
if (process.argv[2]) {
  console.log(syllablize(process.argv[2]))
} else {
  module.exports = syllablize
}
