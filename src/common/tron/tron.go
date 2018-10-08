package tron

import (
	cryp "../crypto"
	"crypto"
	"crypto/ecdsa"
)

func sha3omit12(intput []byte) []byte {
	SHA3_256_h := crypto.SHA3_256.New()
	SHA3_256_h.Reset()
	SHA3_256_h.Write(intput)
	hash := SHA3_256_h.Sum(nil)
	address := hash[10:21]
	address[0] = 0x21
	return address
}

func Pub2Address(p *ecdsa.PublicKey) string {
	pubKey := cryp.FromECDSAPub(32, p)
	address := sha3omit12(pubKey)
	base58 := cryp.B58checkencode(0x00, address)
	return base58
}
