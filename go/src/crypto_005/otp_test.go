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
	e := OTP.Enc(k, m)
	d := OTP.Dec(k, e)
	im, _ := strconv.ParseInt(m, 2, 64)
	id, _ := strconv.ParseInt(d, 2, 64)
	if im != id {
		t.Errorf("Original and decrypted messages are different.")
		return
	}
	// Second, we test if only with the original message
	// and the encrypted message we can get the key.
	ie, _ := strconv.ParseInt(e, 2, 64)
	ik, _ := strconv.ParseInt(k, 2, 64)
	if im ^ ie != ik {
		t.Errorf("It should be possible to compute the key from the original message and the encrypted message.")
		return
	}
}
