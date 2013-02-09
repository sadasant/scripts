package crypto_005_test

import (
	C "crypto_005"
	"strconv"
	"testing"
)

func TestOTP(t *testing.T) {
	m := "0110111"
	k := "1011010"
	e := C.OTP.Enc(k, m)
	d := C.OTP.Dec(k, e)
	im, err := strconv.ParseInt(m, 2, 64)
	if err != nil {
		t.Error(err)
		return
	}
	id, _ := strconv.ParseInt(d, 2, 64)
	if err != nil {
		t.Error(err)
		return
	}
	if im != id {
		t.Errorf("Original and decrypted messages are different.")
	}
}
