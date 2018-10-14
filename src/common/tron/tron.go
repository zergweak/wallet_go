package tron

import (
	cryp "../crypto"
	"crypto"
	"crypto/ecdsa"
	_ "../crypto/sha3"
)

func sha3omit12(input []byte) []byte {
	Sha3256H := crypto.SHA3_256.New()
	Sha3256H.Reset()
	Sha3256H.Write(input[1:])
	hash := Sha3256H.Sum(nil)
	address := hash[11:32]
	address[0] = 0x41
	return address
}

func Pub2Address(p *ecdsa.PublicKey) string {
	pubKey := cryp.FromECDSAPub (p)
	address := sha3omit12(pubKey)
	base58 := cryp.B58checkencode(address)
	return base58
}
