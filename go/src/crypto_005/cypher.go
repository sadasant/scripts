package crypto_005

type Cypher interface {
	Enc(key string, message string) string
	Dec(key string, encmesg string) string
}

type cypher struct {
	message string
	key     string
}
