package crypto_005

type otp cypher

func (o otp) Enc(k byte, m byte) byte {
	return k ^ m
}

func (c otp) Dec(k byte, e byte) byte {
	return k ^ e
}

var OTP otp
