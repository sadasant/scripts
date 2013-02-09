package crypto_005

import "strconv"

type otp cypher

func (o otp) Enc(k string, m string) string {
	ik, _ := strconv.ParseInt(k, 2, 64)
	im, _ := strconv.ParseInt(m, 2, 64)
	c := strconv.FormatInt(ik^im, 2)
	return c
}

func (c otp) Dec(k string, e string) string {
	ik, _ := strconv.ParseInt(k, 2, 64)
	ie, _ := strconv.ParseInt(e, 2, 64)
	m := strconv.FormatInt(ik^ie, 2)
	return m
}

var OTP otp
