package crypto-005

import "strconv"

func OTP_enc (k, m string) string {
	ik, _ := strconv.ParseInt(k, 2, 64)
	im, _ := strconv.ParseInt(m, 2, 64)
	c := strconv.FormatInt(ik ^ im, 2)
	return c
}

func OTP_dec (k, c string) string {
	ik, _ := strconv.ParseInt(k, 2, 64)
	ic, _ := strconv.ParseInt(c, 2, 64)
	m := strconv.FormatInt(ik ^ ic, 2)
	return m
}


func main() {
	m := "0110111"
	k := "1011010"
	println("msg: " + m)
	println("key: " + k)
	c := OTP_enc(k, m)
	println("enc: " + c)
	println("dec: " + OTP_dec(k, c))
}
