// by Daniel Rodriguez @sadasant
// Original code written in C# by [Strilanc](https://github.com/Strilanc)
// Original repo: [NotAMonad](https://github.com/Strilanc/NotAMonad)

// This is the fixed code for an Issue I reported here: https://code.google.com/p/go/issues/detail?can=2&start=0&num=100&q=&colspec=ID%20Status%20Stars%20Priority%20Owner%20Reporter%20Summary&groupby=&sort=&id=6247
// remy_o fixed it with this patch: https://codereview.appspot.com/13216043
//
// Issue:
// I'm trying to run this script in go1.1.2 linux/arm:
//
// http://play.golang.org/p/V-76cJjU3D
//
// It says "reg R13 left allocated" on line 43.
//
// It compiles in play.golang.org, but the calculations doesn't match with the expected value (check lines 85..88), the last element "there" appears with value: complex128=(-0-0.6400000000000001i).
//
// The go env output is:
//
// go env
// GOARCH="arm"
// GOBIN=""
// GOCHAR="5"
// GOEXE=""
// GOHOSTARCH="arm"
// GOHOSTOS="linux"
// GOOS="linux"
// GOPATH="/home/sadasant/code/go"
// GORACE=""
// GOROOT="/usr/lib/go"
// GOTOOLDIR="/usr/lib/go/pkg/tool/linux_arm"
// CC="gcc"
// GOGCCFLAGS="-g -O2 -fPIC -marm -pthread"
// CGO_ENABLED="1"

package main

import (
	"fmt"
	"math"
	"reflect"
)

type Val interface{}
type Superposition map[string]Val

type QuantumSuperposition struct {
	Dict Superposition
}

func (p *QuantumSuperposition) From(dict map[string]Val) QuantumSuperposition {
	p.Dict = dict
	return *p
}

func (p QuantumSuperposition) Transform(t func(string) string) QuantumSuperposition {
	dict := Superposition{}
	for k, v := range p.Dict {
		trans := t(k)
		if _, ok := dict[trans]; ok {
			dict[trans] = dict[trans].(complex128) + v.(complex128)
		} else {
			if _v, ok := v.([]Val); ok {
				dict[trans] = _v[1]
			} else {
				dict[trans] = v
			}
		}
	}
	return new(QuantumSuperposition).From(dict)
}

// FIXME:
// In go1.1.2 linux/arm it breaks with: reg R13 left allocated
func (p QuantumSuperposition) Flatten() QuantumSuperposition {
	dict := Superposition{}
	for _, l := range p.Dict {
		if _l, ok := l.([]Val); ok {
			for k, v := range _l[0].(QuantumSuperposition).Dict {
				if _, ok := dict[k]; ok {
					dict[k] = dict[k].(complex128) + mult(v.(complex128), _l[1].(complex128))
				} else {
					if _v, ok := v.([]Val); ok {
						dict[k] = mult(_v[1].(complex128), _l[1].(complex128))
					} else {
						dict[k] = mult(v.(complex128), _l[1].(complex128))
					}
				}
			}
		} else {
			return p
		}
	}
	return new(QuantumSuperposition).From(dict)
}

// As appears here: http://www.clarku.edu/~djoyce/complex/mult.html
func mult(a, b complex128) complex128 {
	ra := real(a)
	rb := real(b)
	ia := imag(a)
	ib := imag(b)
	r := (ra * rb) - (ia * ib)
	i := (ra * ib) + (rb * ia)
	// Becasue we can have negative zeroes...
	if r == -0 {
		r = 0
	}
	return complex(r, i)
}

func main() {
	a1 := new(QuantumSuperposition).From(Superposition{
		"hey":    complex(1.0/5*3, 0),
		"listen": complex(-1.0/5*4, 0),
	})
	a2 := new(QuantumSuperposition).From(Superposition{
		"over":  complex(1.0/5*3, 0),
		"there": complex(-1.0/5*4, 0),
	})
	s2 := new(QuantumSuperposition).From(Superposition{
		"a1": []Val{a1, complex(1.0/5*3, 0)},
		"a2": []Val{a2, complex(0, 1.0/5*4)},
	})
	s2 = s2.Flatten()
	expected := Superposition{
		"hey":    complex(1.0/25*9, 0),
		"listen": complex(-1.0/25*12, 0),
		"over":   complex(0, 1.0/25*12),
		"there":  complex(0, -1.0/25*16),
	}
	// Fix rounding.
	c := s2.Dict["there"].(complex128)
	i := (math.Floor(imag(c) * 100)/100)+0.01
	s2.Dict["there"] = complex(real(c), i)
	if reflect.DeepEqual(s2.Dict, expected) != true {
		fmt.Printf("DeepEqual\nexpeted: %s\nreceived:%s", expected, s2.Dict)
	}
}
