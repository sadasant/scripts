package crypto_005

import (
	"strconv"
	"testing"
)

func TestOTP(t *testing.T) {
	// First we provide an example message and key,
	// this test is meant to make sure the encrypted message
	// once is decrypted remains consistent with the original message.
	m := "0110111"
	k := "1011010"
	im, _ := strconv.ParseInt(m, 2, 64)
	ik, _ := strconv.ParseInt(k, 2, 64)
	bm := byte(im)
	bk := byte(ik)
	be := OTP.Enc(bk, bm)
	bd := OTP.Dec(bk, be)
	if bm != bd {
		t.Errorf("Original and decrypted messages are different.")
		return
	}
	// Second, we test if only with the original message
	// and the encrypted message we can get the key.
	if bm ^ be != bk {
		t.Errorf("It should be possible to compute the key from the original message and the encrypted message.")
		return
	}
}
