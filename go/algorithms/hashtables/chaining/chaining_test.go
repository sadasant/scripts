package hashtable

import (
	"testing"
)

func TestFill(t *testing.T) {
	h := NewHashTable()
	go_hash := map[string]string{
		"1":  "ONE",
		"2":  "TWO",
		"3":  "THREE",
		"4":  "FOUR",
		"5":  "FIVE",
		"6":  "SIX",
		"7":  "SEVEN",
		"8":  "EIGHT",
		"9":  "NINE",
		"10": "TEN",
	}
	for k, v := range go_hash {
		h.Put(k, v)
	}
	for k, v := range go_hash {
		got := h.Get(k)
		if got != v {
			t.Errorf("Got: %v but wanted %v\n%q", got, v, h)
		}
	}
	h.Del("5")
	if h.Size() != 9 {
		t.Errorf("After h.Del, the size should be 9, not %v\n%q", h.Size, h)
	}
}
